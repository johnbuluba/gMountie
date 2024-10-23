package grpc

import (
	"fmt"
	"gmountie/pkg/server/config"
	"gmountie/pkg/utils/log"
	"net"

	_ "gmountie/pkg/server/grpc/snappy" // Installing the snappy encoding as an available compressor.

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip" // Installing the gzip encoding as an available compressor.
	"google.golang.org/grpc/reflection"
)

// ServiceRegistrar is an interface that defines the ServiceRegistrar method.
type ServiceRegistrar interface {
	Register(*grpc.Server)
}

// Server is a struct that contains a gRPC server.
type Server struct {
	config   *config.Config
	services []ServiceRegistrar
	server   *grpc.Server
}

// NewServer creates a new gRPC server.
func NewServer(config *config.Config, services []ServiceRegistrar) *Server {
	return &Server{
		config:   config,
		services: services,
	}
}

// Serve starts the gRPC server.
func (s *Server) Serve() error {
	// Create a new listener.
	lis, err := s.createListener()
	if err != nil {
		return err
	}
	// Create a new gRPC server.
	s.server = grpc.NewServer(s.getOptions()...)
	// Register the services.
	for _, service := range s.services {
		service.Register(s.server)
	}
	// Add reflection.
	reflection.Register(s.server)
	// Log enabled services.
	for name := range s.server.GetServiceInfo() {
		log.Log.Info("gRPC service is enabled", zap.String("service", name))
	}
	log.Log.Info("gRPC server is running", zap.String("address", lis.Addr().String()))
	// Serve the gRPC server.
	return s.server.Serve(lis)
}

// Stop stops the gRPC server.
func (s *Server) Stop(gracefully bool) {
	if s.server == nil {
		return
	}
	if gracefully {
		s.server.GracefulStop()
	} else {
		s.server.Stop()
	}
}

// createListener creates a new listener.
func (s *Server) createListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", s.config.Server.Address, s.config.Server.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create listener")
	}
	return lis, nil
}

// getOptions returns the gRPC server options.
func (s *Server) getOptions() []grpc.ServerOption {
	unaryLog, streamLog := s.getLoggingInterceptor()
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			unaryLog,
		),
		grpc.ChainStreamInterceptor(
			streamLog,
		),
	}
}

// getLoggingInterceptor returns a new logging interceptor.
func (s *Server) getLoggingInterceptor() (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	unary := logging.UnaryServerInterceptor(InterceptorLogger(log.Log), opts...)
	stream := logging.StreamServerInterceptor(InterceptorLogger(log.Log), opts...)
	return unary, stream
}
