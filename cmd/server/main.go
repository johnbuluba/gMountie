package main

import (
	"flag"
	"grpc-fs/pkg/common"
	"grpc-fs/pkg/server/config"
	"grpc-fs/pkg/server/grpc"
	"grpc-fs/pkg/server/io"
	"os"

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
func Start(cfgString string) error {
	cfg, err := config.LoadConfigFromString(cfgString)
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
	flag.Parse()
	cfg := Config
	if len(flag.Args()) == 1 {
		data, err := os.ReadFile(flag.Arg(0))
		if err != nil {
			common.Log.Sugar().Fatalf("failed to read config file: %v", err)
		}
		cfg = string(data)
	} else {
		common.Log.Info("using default configuration")
	}

	err := Start(cfg)
	if err != nil {
		common.Log.Sugar().Fatalf("failed to start server: %v", err)
	}
}
