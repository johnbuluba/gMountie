package service

import (
	"gmountie/pkg/common"
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/io"

	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"github.com/pkg/errors"
)

// VolumeService is a service that manages volumes.
type VolumeService interface {
	// List lists all volumes.
	List() ([]common.Volume, error)
	// GetVolumeFileSystem gets the filesystem for a volume.
	GetVolumeFileSystem(name string) (pathfs.FileSystem, error)
}

type VolumeServiceOptions func(*VolumeServiceImpl)

func WithMiddleware(middleware ...io.Middleware) VolumeServiceOptions {
	return func(s *VolumeServiceImpl) {
		s.middleware = append(s.middleware, middleware...)
	}
}

// VolumeServiceImpl is an implementation of the VolumeService interface.
type VolumeServiceImpl struct {
	config      *config.Config
	filesystems map[string]pathfs.FileSystem
	middleware  []io.Middleware
}

// NewVolumeService creates a new VolumeService.
func NewVolumeService(cfg *config.Config, options ...VolumeServiceOptions) VolumeService {
	fs := make(map[string]pathfs.FileSystem)
	svc := &VolumeServiceImpl{
		config:      cfg,
		filesystems: fs,
		middleware:  make([]io.Middleware, 0),
	}
	for _, option := range options {
		option(svc)
	}
	for _, v := range cfg.Volumes {
		svc.addFileSystem(v.Name, io.NewLocalFilesystem(v.Path))
	}
	return svc
}

// List lists all volumes.
func (s *VolumeServiceImpl) List() ([]common.Volume, error) {
	volumes := make([]common.Volume, 0)
	for _, v := range s.config.Volumes {
		volumes = append(volumes, common.Volume{Name: v.Name})
	}
	return volumes, nil
}

// GetVolumeFileSystem gets the filesystem for a volume.
func (s *VolumeServiceImpl) GetVolumeFileSystem(name string) (pathfs.FileSystem, error) {
	fs, ok := s.filesystems[name]
	if !ok {
		return nil, errors.Errorf("volume %s not found", name)
	}
	return fs, nil
}

// addFileSystem adds a filesystem to the volume service.
func (s *VolumeServiceImpl) addFileSystem(name string, fs pathfs.FileSystem) {
	// Apply middleware
	for _, currentMiddleware := range s.middleware {
		fs = currentMiddleware(fs)
	}
	s.filesystems[name] = fs
}
