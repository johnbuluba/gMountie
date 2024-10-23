// Copyright 2016 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is main program driver for the loopback filesystem from
// github.com/hanwen/go-fuse/fs/, a filesystem that shunts operations
// to an underlying file system.
package main

import (
	"flag"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/client/service"
	"gmountie/pkg/utils/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Log.Sugar().Fatalf("Usage:\n  hello MOUNTPOINT")
	}

	//client, err := grpc.NewClient("192.168.11.42:9449", "data")
	//client, err := grpc.NewClient("gmountie.home.buluba.net:443", "data")
	client, err := grpc.NewClient("localhost:9449")
	if err != nil {
		log.Log.Sugar().Fatalf("failed to create client: %v", err)
	}
	client.Connect()

	mounter := service.NewMounterService(client)
	err = mounter.Mount("test", flag.Arg(0))
	if err != nil {
		log.Log.Sugar().Fatalf("failed to mount: %v", err)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	err = mounter.UnmountAll()
	if err != nil {
		log.Log.Sugar().Fatalf("failed to unmount: %v", err)
		return
	}

}
