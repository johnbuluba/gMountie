package utils

import (
	"context"
	"gmountie/pkg/client"
	grpcClient "gmountie/pkg/client/grpc"
	"gmountie/pkg/server"
	"gmountie/pkg/server/config"
	grpcServer "gmountie/pkg/server/grpc"
	"net"

	"github.com/thanhpk/randstr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// AppTestingContext is a struct that holds the testing context
type AppTestingContext struct {
	// cfg is the configuration for the server.
	cfg config.Config
	// serverCtx is the application context.
	serverCtx *server.AppContext
	// clientCtx is the client context.
	clientCtx *client.AppContext
	// clientOptions is the client options.
	clientOptions []grpcClient.ClientOption
	// serverOptions is the server options.
	serverOptions []grpcServer.ServerOption
	// listener is the listener for the gRPC server.
	listener *bufconn.Listener
	// server is the gRPC server.
	server *grpcServer.Server
	// client is the gRPC client.
	client *grpcClient.Client
	// volumes are the test volumes.
	volumes []*TestVolume
}

// TestOptions is a type that defines the TestOptions function.
type TestOptions func(*AppTestingContext)

// WithBasicAuth sets the basic authentication for the testing context.
func WithBasicAuth(username, password string) TestOptions {
	return func(c *AppTestingContext) {
		// Set the server basic auth
		c.cfg.Auth = &config.BasicAuthConfig{
			Users: []config.BasicAuthConfigUser{
				{
					Username: username, Password: password,
				},
			},
		}
		// Append the client options
		c.clientOptions = append(c.clientOptions, grpcClient.WithBasicAuth(username, password))
	}
}

// WithRandomTestVolume creates random test volume.
func WithRandomTestVolume(randomfiles bool) TestOptions {
	return func(c *AppTestingContext) {
		v, err := NewTestVolume(randstr.String(10), randomfiles)
		if err != nil {
			panic(err)
		}
		c.volumes = append(c.volumes, v)
		// Add in server config.
		c.cfg.Volumes = append(c.cfg.Volumes, &config.VolumeConfig{
			Name: v.Name,
			Path: v.GetSrcPath(),
		})
	}
}

// NewAppTestingContext creates a new AppTestingContext.
func NewAppTestingContext(options ...TestOptions) (*AppTestingContext, error) {
	appCtx := &AppTestingContext{}
	appCtx.cfg.Server = &config.ServerConfig{Metrics: false}
	// Apply the options
	for _, opt := range options {
		opt(appCtx)
	}
	// Create a new server app context
	appCtx.serverCtx = server.NewServerAppContext(&appCtx.cfg)
	// Create listener
	appCtx.listener = bufconn.Listen(1024 * 1024)
	// Create a new server
	appCtx.serverOptions = append(appCtx.serverOptions, grpcServer.WithListener(appCtx.listener))
	appCtx.server = grpcServer.NewServer(
		&config.Config{},
		appCtx.serverCtx.AuthService,
		appCtx.serverCtx.GetGrpcServices(),
		appCtx.serverOptions...,
	)
	// Create a new client
	appCtx.clientOptions = append(appCtx.clientOptions, grpcClient.WithDialOptions([]grpc.DialOption{
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return appCtx.listener.Dial()
		}),
	}))
	c, err := grpcClient.NewClient("passthrough://bufnet", appCtx.clientOptions...)
	if err != nil {
		return nil, err
	}
	appCtx.client = c
	appCtx.clientCtx = client.NewAppContext(c)
	return appCtx, nil
}

// GetServerApp returns the server app context.
func (c *AppTestingContext) GetServerApp() *server.AppContext {
	return c.serverCtx
}

// GetClientApp returns the client app context.
func (c *AppTestingContext) GetClientApp() *client.AppContext {
	return c.clientCtx
}

// GetClient returns the gRPC client.
func (c *AppTestingContext) GetClient() *grpcClient.Client {
	return c.client
}

// GetVolumes returns the test volumes.
func (c *AppTestingContext) GetVolumes() []*TestVolume {
	return c.volumes
}

// MountVolume mounts the test volume.
func (c *AppTestingContext) MountVolume(v *TestVolume) error {
	return c.clientCtx.MounterService.Mount(v.Name, v.GetMountPath())
}

// Start starts the gRPC server.
func (c *AppTestingContext) Start() error {
	go func() {
		if err := c.server.Serve(); err != nil {
			panic(err)
		}
	}()
	return nil
}

// Close closes the gRPC server.
func (c *AppTestingContext) Close() error {
	c.server.Stop(false)
	// Close the volumes
	for _, v := range c.volumes {
		if err := v.Close(); err != nil {
			return err
		}
	}
	return nil
}
