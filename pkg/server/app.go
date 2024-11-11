package server

import (
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/controller"
	"gmountie/pkg/server/grpc"
	"gmountie/pkg/server/io"
	"gmountie/pkg/server/io/middleware"
	"gmountie/pkg/server/service"
	"gmountie/pkg/utils/log"
	"runtime"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AppContext struct {
	// Config is the configuration for the server.
	Config        *config.Config
	VolumeService service.VolumeService
	AuthService   service.AuthService
}

// NewServerAppContext creates a new ServerContext.
func NewServerAppContext(cfg *config.Config) *AppContext {
	volumeService := service.NewVolumeService(cfg, service.WithMiddleware(getVolumeMiddleware()...))
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

// getVolumeMiddleware returns the volume middleware.
func getVolumeMiddleware() []io.Middleware {
	m := make([]io.Middleware, 0)
	// If user is root we can assume the user identity
	if runtime.GOOS == "linux" && syscall.Getuid() == 0 {
		m = append(m, middleware.AssumeUserMiddleware)
	}
	// Print middleware
	names := make([]string, 0, len(m))
	for _, mw := range m {
		names = append(names, mw.GetName())
	}
	if len(names) > 0 {
		log.Log.Info("enabled filesystem middlewares", zap.Strings("middlewares", names))
	}
	return m
}
