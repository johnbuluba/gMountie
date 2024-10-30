package client

import (
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/service"
	"gmountie/pkg/utils/log"
)

// AppContext is a struct that holds the application context.
type AppContext struct {
	client         *grpc.Client
	MounterService service.MounterService
	VolumeService  service.VolumeService
}

// NewAppContext creates a new AppContext.
func NewAppContext(client *grpc.Client) *AppContext {
	log.Log.Info("creating app context")
	return &AppContext{
		client:         client,
		MounterService: service.NewMounterService(client),
		VolumeService:  service.NewVolumeService(client),
	}
}

// Close closes the AppContext
func (a *AppContext) Close() error {
	log.Log.Info("closing app context")
	err := a.MounterService.UnmountAll()
	if err != nil {
		return err
	}
	return a.client.Close()
}
