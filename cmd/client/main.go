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
	"gmountie/pkg/client/io"
	"gmountie/pkg/common"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"go.uber.org/zap"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  hello MOUNTPOINT")
	}

	//client, err := grpc.NewClient("192.168.11.42:9449", "data")
	//client, err := grpc.NewClient("gmountie.home.buluba.net:443", "data")
	client, err := grpc.NewClient("localhost:9449", "test")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fs := io.NewGrpcInode(client)
	fs.SetDebug(true)
	nodeFS := pathfs.NewPathNodeFs(fs, &pathfs.PathNodeFsOptions{ClientInodes: true, Debug: true})
	opts := nodefs.NewOptions()
	opts.Debug = true
	sec := time.Second
	connector := nodefs.NewFileSystemConnector(nodeFS.Root(),
		&nodefs.Options{
			EntryTimeout:        sec,
			AttrTimeout:         sec,
			NegativeTimeout:     0.0,
			Debug:               true,
			LookupKnownChildren: true,
		})
	zap.NewStdLog(common.Log)
	server, err := fuse.NewServer(
		connector.RawFS(), flag.Arg(0), &fuse.MountOptions{
			AllowOther:     true,
			SingleThreaded: false,
			Debug:          true,
			Logger:         zap.NewStdLog(common.Log),
		})
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		errRetry := retry.Do(
			func() error {
				err = server.Unmount()
				if err != nil {
					common.Log.Sugar().Errorf("Unmount fail: %v", err)
					return err
				}
				return nil
			},
			retry.Attempts(3),
			retry.Delay(5*time.Second),
		)
		if errRetry != nil {
			common.Log.Sugar().Fatalf("Unmount fail, giving up: %v\n", errRetry)
		}
	}()
	server.Serve()
	server.Wait()
}
