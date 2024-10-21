package server

import (
	"context"
	"errors"
	"grpc-fs/pkg/common"
	"grpc-fs/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"go.uber.org/zap"
)

type fileEntry struct {
	file nodefs.File
	path string
}

type RpcFileServerImpl struct {
	filesystem pathfs.FileSystem
	files      map[string]*fileEntry
	proto.UnimplementedRpcFileServer
}

func NewRpcFileServer(filesystem pathfs.FileSystem) *RpcFileServerImpl {
	return &RpcFileServerImpl{
		filesystem: filesystem,
		files:      make(map[string]*fileEntry),
	}
}

func (r *RpcFileServerImpl) Open(ctx context.Context, request *proto.OpenRequest) (*proto.OpenReply, error) {
	common.Log.Debug("rpc: Open", zap.String("path", request.Path), zap.Uint32("flags", request.Flags))
	fd, status := r.filesystem.Open(request.Path, request.Flags, createContext(ctx))
	reply := &proto.OpenReply{
		Status: int32(status),
	}
	r.addFile(request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateReply, error) {
	common.Log.Debug("rpc: Create",
		zap.String("path", request.Path),
		zap.Uint32("flags", request.Flags),
		zap.Uint32("mode", request.Mode),
	)
	fd, status := r.filesystem.Create(request.Path, request.Flags, request.Mode, createContext(ctx))
	reply := &proto.CreateReply{
		Status: int32(status),
	}
	r.addFile(request.Path, fd)
	return reply, nil
}

func (r *RpcFileServerImpl) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadReply, error) {
	common.Log.Debug("rpc: Read",
		zap.String("path", request.Path),
		zap.Uint32("size", request.Size),
		zap.Int64("offset", request.Offset),
	)
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, request.Size)
	n, status := file.file.Read(buf, request.Offset)
	if status != fuse.OK {
		return &proto.ReadReply{
			Status: int32(status),
		}, nil
	}
	buf, status = n.Bytes(buf)
	reply := &proto.ReadReply{
		Size:   int64(n.Size()),
		Bytes:  buf,
		Status: int32(status),
	}
	return reply, nil
}

func (r *RpcFileServerImpl) Write(ctx context.Context, request *proto.WriteRequest) (*proto.WriteReply, error) {
	common.Log.Debug("rpc: Write",
		zap.String("path", request.Path),
		zap.Int64("offset", request.Offset),
	)
	file, err := r.getFile(request.Path)
	if err != nil {
		return nil, err
	}
	written, status := file.file.Write(request.Bytes, request.Offset)

	reply := &proto.WriteReply{
		Written: written,
		Status:  int32(status),
	}
	return reply, nil
}

func (r *RpcFileServerImpl) addFile(path string, file nodefs.File) {
	r.files[path] = &fileEntry{
		file: file,
		path: path,
	}
}

func (r *RpcFileServerImpl) getFile(path string) (*fileEntry, error) {
	_, ok := r.files[path]
	if !ok {
		return nil, errors.New("file not found")
	}
	return r.files[path], nil
}
