package client

import (
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/service"
)

// AppContext is a struct that holds the application context.
type AppContext struct {
	MounterService service.MounterService
	VolumeService  service.VolumeService
}

// NewAppContext creates a new AppContext.
func NewAppContext(client *grpc.Client) *AppContext {

	return &AppContext{
		MounterService: service.NewMounterService(client),
		VolumeService:  service.NewVolumeService(client),
	}
}
