package io

import (
	"context"
	"gmountie/pkg/common"
	"gmountie/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"go.uber.org/zap"
)

type GrpcFile struct {
	fileClient proto.RpcFileClient
	path       string
	nodefs.File
}

func NewGrpcFile(fileClient proto.RpcFileClient, path string) *GrpcFile {
	return &GrpcFile{
		fileClient: fileClient,
		path:       path,
		File:       nodefs.NewDefaultFile(),
	}
}

func (f *GrpcFile) Read(dest []byte, off int64) (fuse.ReadResult, fuse.Status) {
	res, err := f.fileClient.Read(context.Background(), &proto.ReadRequest{
		Path:   f.path,
		Offset: off,
		Size:   uint32(len(dest)),
	})
	if err != nil || res == nil {
		common.Log.Error("error in call: Read", zap.String("path", f.path), zap.Error(err))
		return nil, fuse.EIO
	}
	return fuse.ReadResultData(res.Bytes), fuse.Status(res.Status)
}
func (f *GrpcFile) Write(data []byte, off int64) (written uint32, code fuse.Status) {
	res, err := f.fileClient.Write(context.Background(), &proto.WriteRequest{
		Path:   f.path,
		Offset: off,
		Bytes:  data,
	})
	if err != nil || res == nil {
		common.Log.Error("error in call: Write", zap.String("path", f.path), zap.Error(err))
		return 0, fuse.EIO
	}
	return res.Written, fuse.Status(res.Status)
}
