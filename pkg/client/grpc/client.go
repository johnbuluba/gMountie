package grpc

import (
	"grpc-fs/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (proto.RpcFsClient, proto.RpcFileClient, error) {
	conn, err := grpc.NewClient("localhost:9449", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return proto.NewRpcFsClient(conn), proto.NewRpcFileClient(conn), nil
}
