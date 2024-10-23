package grpc

import (
	"gmountie/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a struct that holds the gRPC client
type Client struct {
	volume string
	conn   *grpc.ClientConn
	Fs     proto.RpcFsClient
	File   proto.RpcFileClient
}

func NewClient(endpoint string, volume string) (*Client, error) {
	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithChainUnaryInterceptor(
			getInterceptors(volume)...,
		),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:   conn,
		volume: volume,
		Fs:     proto.NewRpcFsClient(conn),
		File:   proto.NewRpcFileClient(conn),
	}, nil
}

// GetInterceptors returns the client interceptors
func getInterceptors(volume string) []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{
		NewVolumeInterceptor(volume),
	}
}
