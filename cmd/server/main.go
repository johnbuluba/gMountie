package main

import (
	"flag"
	"gmountie/pkg/common"
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/controller"
	"gmountie/pkg/server/grpc"
	"gmountie/pkg/server/service"
	"os"

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

	volumeService := service.NewVolumeService(&cfg)
	server := grpc.NewServer(&cfg, []grpc.ServiceRegistrar{
		controller.NewGrpcServer(volumeService),
		controller.NewRpcFileServer(volumeService),
		controller.NewVolumeService(volumeService),
	})

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
