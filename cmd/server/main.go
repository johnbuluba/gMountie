package main

import (
	"flag"
	"gmountie/pkg/server"
	"gmountie/pkg/server/config"
	"gmountie/pkg/server/grpc"
	"gmountie/pkg/utils/log"
	"os"

	"github.com/pkg/errors"
)

const Config = `
server:
  address: 127.0.0.1
auth:
  type: basic
  users:
    - username: john
      password: 123456
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

	context := server.NewServerAppContext(&cfg)

	s := grpc.NewServer(&cfg, context.AuthService, context.GetGrpcServices())

	if err = s.Serve(); err != nil {
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
			log.Log.Sugar().Fatalf("failed to read config file: %v", err)
		}
		cfg = string(data)
	} else {
		log.Log.Info("using default configuration")
	}

	err := Start(cfg)
	if err != nil {
		log.Log.Sugar().Fatalf("failed to start server: %v", err)
	}
}
