package io

import (
	"context"
	"gmountie/pkg/common"
	"gmountie/pkg/proto"
	"gmountie/pkg/server/grpc/snappy"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	},
		grpc.UseCompressor(snappy.Name),
	)
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
	},
		grpc.UseCompressor(snappy.Name),
	)
	if err != nil || res == nil {
		common.Log.Error("error in call: Write", zap.String("path", f.path), zap.Error(err))
		return 0, fuse.EIO
	}
	return res.Written, fuse.Status(res.Status)
}

// Release is called when the file is closed
func (f *GrpcFile) Release() {
	_, err := f.fileClient.Release(context.Background(), &proto.ReleaseRequest{Path: f.path})
	if err != nil {
		common.Log.Error("error in call: Release", zap.String("path", f.path), zap.Error(err))
	}
}

// Flush is called when the file is closed
func (f *GrpcFile) Flush() fuse.Status {
	res, err := f.fileClient.Flush(context.Background(), &proto.FlushRequest{Path: f.path})
	if err != nil {
		common.Log.Error("error in call: Flush", zap.String("path", f.path), zap.Error(err))
	}
	return fuse.Status(res.Status)
}

// Fsync is called to flush the file data
func (f *GrpcFile) Fsync(flags int) fuse.Status {
	res, err := f.fileClient.Fsync(context.Background(), &proto.FsyncRequest{Path: f.path, Flags: int64(flags)})
	if err != nil {
		common.Log.Error("error in call: Fsync", zap.String("path", f.path), zap.Error(err))
	}
	return fuse.Status(res.Status)
}

// ----- File locking

// GetLk is called to retrieve the lock information
func (f *GrpcFile) GetLk(owner uint64, lk *fuse.FileLock, flags uint32, out *fuse.FileLock) fuse.Status {
	res, err := f.fileClient.GetLk(context.Background(), &proto.GetLkRequest{
		Path:  f.path,
		Owner: owner,
		Flags: flags,
		Lk:    &proto.FileLock{Start: lk.Start, End: lk.End, Typ: lk.Typ, Pid: lk.Pid},
	})
	if err != nil {
		common.Log.Error("error in call: GetLk", zap.String("path", f.path), zap.Error(err))
	}
	out = &fuse.FileLock{Start: res.Lk.Start, End: res.Lk.End, Typ: res.Lk.Typ, Pid: res.Lk.Pid}
	return fuse.Status(res.Status)
}

// SetLk is called to set the lock information
func (f *GrpcFile) SetLk(owner uint64, lk *fuse.FileLock, flags uint32) fuse.Status {
	res, err := f.fileClient.SetLk(context.Background(), &proto.SetLkRequest{
		Path:  f.path,
		Owner: owner,
		Flags: flags,
		Lk:    &proto.FileLock{Start: lk.Start, End: lk.End, Typ: lk.Typ, Pid: lk.Pid},
	})
	if err != nil {
		common.Log.Error("error in call: SetLk", zap.String("path", f.path), zap.Error(err))
	}
	return fuse.Status(res.Status)
}

// SetLkw is called to set the lock information
func (f *GrpcFile) SetLkw(owner uint64, lk *fuse.FileLock, flags uint32) fuse.Status {
	res, err := f.fileClient.SetLkw(context.Background(), &proto.SetLkwRequest{
		Path:  f.path,
		Owner: owner,
		Flags: flags,
		Lk:    &proto.FileLock{Start: lk.Start, End: lk.End, Typ: lk.Typ, Pid: lk.Pid},
	})
	if err != nil {
		common.Log.Error("error in call: SetLkw", zap.String("path", f.path), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}
