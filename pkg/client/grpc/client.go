package grpc

import (
	"gmountie/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client is a struct that holds the gRPC client
type Client struct {
	conn   *grpc.ClientConn
	Fs     proto.RpcFsClient
	File   proto.RpcFileClient
	Volume proto.VolumeServiceClient
}

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithChainUnaryInterceptor(
			getInterceptors()...,
		),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:   conn,
		Fs:     proto.NewRpcFsClient(conn),
		File:   proto.NewRpcFileClient(conn),
		Volume: proto.NewVolumeServiceClient(conn),
	}, nil
}

// GetInterceptors returns the client interceptors
func getInterceptors() []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{}
}
