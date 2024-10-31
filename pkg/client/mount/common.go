package mount

import (
	"fmt"
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
	debug      = true
)

// Mounter is the interface for the mounter
type Mounter interface {
	// IsVolumeMounted checks if a volume is mounted
	IsVolumeMounted(volume string) bool
	// Unmount unmounts a volume
	Unmount(volumeName string) error
	// UnmountAll unmounts all volumes from the root filesystem
	UnmountAll() error
	// Close closes the root filesystem
	Close() error
}

// ---------------------- Utils ----------------------

// createMountOptions returns the mount options
func createMountOptions(endpoint, volume string) *fuse.MountOptions {
	return &fuse.MountOptions{
		AllowOther:     false,
		SingleThreaded: false,
		Debug:          debug,
		EnableLocks:    true,
		Name:           FuseFSName,
		FsName:         fmt.Sprintf("%s://%s/%s", FuseFSName, endpoint, volume),
		Logger:         zap.NewStdLog(log.Log.Named("fuse")),
	}
}

// createConnectorOptions returns the connector options
func createConnectorOptions() *nodefs.Options {
	sec := time.Second
	return &nodefs.Options{
		EntryTimeout:        sec,
		AttrTimeout:         sec,
		NegativeTimeout:     0.0,
		Debug:               debug,
		LookupKnownChildren: true,
	}
}

// createFsOptions returns the filesystem options
func createFsOptions() *pathfs.PathNodeFsOptions {
	return &pathfs.PathNodeFsOptions{
		ClientInodes: false,
		Debug:        debug,
	}
}

// stopServer stops the server
func stopServer(server *fuse.Server) error {
	return retry.Do(
		func() error {
			err := server.Unmount()
			if err != nil {
				log.Log.Warn("unmount fail, retrying ...", zap.Error(err))
				return err
			}
			return nil
		},
		retry.Attempts(3),
		retry.Delay(5*time.Second),
	)
}
