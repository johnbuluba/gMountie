package client

import (
	"errors"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/mount"
	"gmountie/pkg/client/service"
	"gmountie/pkg/utils/log"
)

// AppContext is a struct that holds the application context.
type AppContext struct {
	client              grpc.Client
	VolumeService       service.VolumeService
	SingleVolumeMounter mount.SingleVolumeMounter
	MultiVolumeMounter  mount.VFSVolumeMounter
}

// NewAppContext creates a new AppContext.
func NewAppContext(client grpc.Client, multiMountPath string) *AppContext {
	log.Log.Info("creating app context")
	return &AppContext{
		client:              client,
		SingleVolumeMounter: mount.NewSingleVolumeMounter(client),
		MultiVolumeMounter:  mount.NewMultiVolumeMounter(client, multiMountPath),
		VolumeService:       service.NewVolumeService(client),
	}
}

// Close closes the AppContext
func (a *AppContext) Close() error {
	log.Log.Info("closing app context")
	var errs = make([]error, 0)
	errs = append(errs, a.SingleVolumeMounter.Close())
	errs = append(errs, a.MultiVolumeMounter.Close())
	errs = append(errs, a.client.Close())
	return errors.Join(errs...)
}
