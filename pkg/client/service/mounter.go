package service

import (
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/io"
	"gmountie/pkg/utils/log"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"go.uber.org/zap"
)

const (
	FuseFSName = "gMountie"
)

type mount struct {
	// Volume is the volume to mount
	volume    string
	path      string
	server    *fuse.Server
	connector *nodefs.FileSystemConnector
}

// Serve starts the mount
func (m *mount) Serve() {
	log.Log.Info("mounting volume", zap.String("volume", m.volume), zap.String("path", m.path))
	m.server.Serve()
	m.server.Wait()
}

// Wait waits for the mount to finish
func (m *mount) Wait() error {
	return m.server.WaitMount()
}

// Stop stops the mount
func (m *mount) Stop() error {
	log.Log.Info("unmounting volume", zap.String("volume", m.volume), zap.String("path", m.path))
	errRetry := retry.Do(
		func() error {
			err := m.server.Unmount()
			if err != nil {
				log.Log.Sugar().Errorf("unmount fail: %v", err)
				return err
			}
			return nil
		},
		retry.Attempts(3),
		retry.Delay(5*time.Second),
	)
	if errRetry != nil {
		log.Log.Sugar().Fatalf("unmount fail, giving up: %v\n", errRetry)
		return errRetry
	}
	return nil
}

// MounterService is an interface that mounts volumes
type MounterService interface {
	// Mount mounts a volume
	Mount(volume, path string) error
	// Unmount unmounts a volume
	Unmount(volume string) error
	// UnmountAll unmounts all volumes
	UnmountAll() error
}

// MounterServiceImpl is a service that mounts volumes
type MounterServiceImpl struct {
	client *grpc.Client
	mounts map[string]*mount
}

// NewMounterService creates a new MounterServiceImpl
func NewMounterService(client *grpc.Client) MounterService {
	return &MounterServiceImpl{
		client: client,
		mounts: make(map[string]*mount),
	}
}

// Mount mounts a volume
func (m *MounterServiceImpl) Mount(volume, path string) error {
	fs := io.NewLocalFileSystem(m.client, volume)
	fs.SetDebug(true)
	nodeFS := pathfs.NewPathNodeFs(fs, &pathfs.PathNodeFsOptions{
		ClientInodes: true,
		Debug:        true,
	})
	sec := time.Second
	connector := nodefs.NewFileSystemConnector(nodeFS.Root(),
		&nodefs.Options{
			EntryTimeout:        sec,
			AttrTimeout:         sec,
			NegativeTimeout:     0.0,
			Debug:               true,
			LookupKnownChildren: true,
		})
	server, err := fuse.NewServer(
		connector.RawFS(),
		path,
		&fuse.MountOptions{
			AllowOther:     false,
			SingleThreaded: false,
			Debug:          true,
			Name:           FuseFSName,
			FsName:         FuseFSName + ":/" + volume,
			Logger:         zap.NewStdLog(log.Log),
		},
	)
	if err != nil {
		log.Log.Sugar().Fatalf("mount fail: %v\n", err)
	}
	mount := &mount{
		volume:    volume,
		path:      path,
		server:    server,
		connector: connector,
	}
	m.mounts[volume] = mount
	go func() {
		mount.Serve()
	}()
	return mount.Wait()
}

// Unmount unmounts a volume
func (m *MounterServiceImpl) Unmount(volume string) error {
	mount, ok := m.mounts[volume]
	if !ok {
		return nil
	}
	err := mount.Stop()
	if err != nil {
		return err
	}
	delete(m.mounts, volume)
	return nil
}

// UnmountAll unmounts all volumes
func (m *MounterServiceImpl) UnmountAll() error {
	for volume := range m.mounts {
		err := m.Unmount(volume)
		if err != nil {
			return err
		}
	}
	return nil
}
