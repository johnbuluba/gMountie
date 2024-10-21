package main

import (
	"grpc-fs/pkg/server/config"
	"grpc-fs/pkg/server/grpc"
	"grpc-fs/pkg/server/io"
	"log"

	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"github.com/pkg/errors"
)

const Config = `
server:
  address: 127.0.0.1
volumes:
- name: test
  path: /home/john/mnt/test
`

// Start the gRPC server
func Start() error {
	cfg, err := config.LoadConfigFromString(Config)
	if err != nil {
		return err
	}

	fs := pathfs.NewLoopbackFileSystem(cfg.Volumes[0].Path)
	//fs = pathfs.NewLockingFileSystem(fs)
	server := grpc.NewServer(&cfg, []grpc.ServiceRegistrar{io.NewGrpcServer(fs), io.NewRpcFileServer(fs)})

	if err = server.Serve(); err != nil {
		return errors.Wrap(err, "failed to start server")
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
