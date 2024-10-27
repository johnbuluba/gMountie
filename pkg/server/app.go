package server

import (
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/controller"
	"gmountie/pkg/server/grpc"
	"gmountie/pkg/server/service"
)

type AppContext struct {
	// Config is the configuration for the server.
	Config        *config.Config
	VolumeService service.VolumeService
	AuthService   service.AuthService
}

// NewServerAppContext creates a new ServerContext.
func NewServerAppContext(cfg *config.Config) *AppContext {
	volumeService := service.NewVolumeService(cfg)
	authService := service.NewAuthServiceFromConfig(cfg.Auth)
	return &AppContext{
		Config:        cfg,
		VolumeService: volumeService,
		AuthService:   authService,
	}
}

// GetGrpcServices returns the gRPC services.
func (c *AppContext) GetGrpcServices() []grpc.ServiceRegistrar {
	return []grpc.ServiceRegistrar{
		controller.NewGrpcServer(c.VolumeService),
		controller.NewRpcFileServer(c.VolumeService),
		controller.NewVolumeService(c.VolumeService),
	}
}
