package server

import (
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/controller"
	"gmountie/pkg/server/grpc"
	"gmountie/pkg/server/service"

	"github.com/pkg/errors"
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

// Start starts the server.
func Start(cfg *config.Config) error {
	context := NewServerAppContext(cfg)

	s := grpc.NewServer(
		cfg,
		context.AuthService,
		context.GetGrpcServices(),
	)

	if err := s.Serve(); err != nil {
		return errors.Wrap(err, "failed to start server")
	}
	return nil
}
