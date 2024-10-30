package grpc

import (
	"gmountie/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a struct that holds the gRPC client
type Client struct {
	conn        *grpc.ClientConn
	dialOptions []grpc.DialOption
	Fs          proto.RpcFsClient
	File        proto.RpcFileClient
	Volume      proto.VolumeServiceClient
}

// -------------------- Client Options --------------------

// ClientOption is a type that defines the ClientOption function
type ClientOption func(*Client)

// WithDialOptions sets the dial options for the gRPC client
func WithDialOptions(dialOptions []grpc.DialOption) ClientOption {
	return func(c *Client) {
		// Append the dial options
		c.dialOptions = append(c.dialOptions, dialOptions...)
	}
}

// WithBasicAuth sets the basic authentication for the gRPC client
func WithBasicAuth(username, password string) ClientOption {
	return func(c *Client) {
		c.dialOptions = append(c.dialOptions, grpc.WithPerRPCCredentials(NewBasicAuthCredentials(username, password)))
	}
}

// ---------------------- Constructor ----------------------

// NewClient creates a new gRPC client
func NewClient(endpoint string, options ...ClientOption) (*Client, error) {
	c := Client{}
	for _, opt := range options {
		opt(&c)
	}
	conn, err := grpc.NewClient(
		endpoint,
		c.getDialOptions()...,
	)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	c.File = proto.NewRpcFileClient(conn)
	c.Fs = proto.NewRpcFsClient(conn)
	c.Volume = proto.NewVolumeServiceClient(conn)
	return &c, nil
}

// ---------------------- Methods -----------------------

// Connect connects to the gRPC server
func (c *Client) Connect() {
	c.conn.Connect()
}

// Close closes the gRPC client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// GetInterceptors returns the client interceptors
func getInterceptors() []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{
		//logging.UnaryClientInterceptor(grpc2.InterceptorLogger(log.Log)),
	}
}

// getDialOptions returns the dial options
func (c *Client) getDialOptions() []grpc.DialOption {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithChainUnaryInterceptor(
			getInterceptors()...,
		),
	}

	// Append the dial options
	opts = append(opts, c.dialOptions...)
	return opts
}
