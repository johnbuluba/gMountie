// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is main program driver for the loopback filesystem from
// github.com/hanwen/go-fuse/fs/, a filesystem that shunts operations
// to an underlying file system.
package main

import (
	"flag"
	"gmountie/pkg/client"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/utils/log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Log.Sugar().Fatalf("Usage:\n  hello MOUNTPOINT")
	}

	//c, err := grpc.NewClient("192.168.11.42:9449", "data")
	//c, err := grpc.NewClient("gmountie.home.buluba.net:443", "data")
	//c, err := grpc.NewClient("18.194.209.199:9449", grpc.WithBasicAuth("john", "123456"))
	c, err := grpc.NewClient("127.0.0.1:9449", grpc.WithBasicAuth("john", "123456"))
	if err != nil {
		log.Log.Sugar().Fatalf("failed to create client: %v", err)
	}
	c.Connect()

	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Log.Sugar().Fatalf("failed to get home directory: %v", err)
	}
	path := filepath.Join(homePath, "mnt", "gmountie")
	appCtx := client.NewAppContext(c, path)

	defer func() {
		err = appCtx.Close()
		if err != nil {
			log.Log.Sugar().Fatalf("failed to close the application: %v", err)
			return
		}
	}()
	err = appCtx.MultiVolumeMounter.Start()
	if err != nil {
		log.Log.Sugar().Fatalf("failed to start the multi volume mounter: %v", err)
		return
	}
	err = appCtx.MultiVolumeMounter.Mount("test1")
	if err != nil {
		log.Log.Sugar().Fatalf("failed to mount: %v", err)
		return
	}
	err = appCtx.MultiVolumeMounter.Mount("test2")
	if err != nil {
		log.Log.Sugar().Fatalf("failed to mount: %v", err)
		return
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
