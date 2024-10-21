package main

import (
	"fmt"
	"grpc-fs/pkg/server"
	"log"
	"net"

	pb "grpc-fs/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Start the gRPC server
func Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8085))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	if err != nil {
		return err
	}
	fs := pathfs.NewLoopbackFileSystem("/home/john/mnt/test")
	fs = pathfs.NewLockingFileSystem(fs)
	pb.RegisterRpcFsServer(grpcServer, server.NewGrpcServer(fs))
	pb.RegisterRpcFileServer(grpcServer, server.NewRpcFileServer(fs))
	reflection.Register(grpcServer)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

// Main function
func main() {
	err := Start()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
