package io

import (
	"context"
	"grpc-fs/pkg/proto"
	"os"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"google.golang.org/grpc"
)

type RpcServerImpl struct {
	filesystem pathfs.FileSystem
	proto.UnimplementedRpcFsServer
}

// NewGrpcServer creates a new gRPC server
func NewGrpcServer(filesystem pathfs.FileSystem) *RpcServerImpl {
	return &RpcServerImpl{
		filesystem: filesystem,
	}
}

// Register registers the gRPC server
func (r *RpcServerImpl) Register(server *grpc.Server) {
	proto.RegisterRpcFsServer(server, r)
}

func (r *RpcServerImpl) GetAttr(ctx context.Context, request *proto.GetAttrRequest) (*proto.GetAttrReply, error) {
	attr, status := r.filesystem.GetAttr(request.Path, createContext(ctx))
	if attr == nil {
		return &proto.GetAttrReply{
			Status: int32(status),
		}, nil
	}
	reply := &proto.GetAttrReply{
		Attributes: &proto.Attr{
			Ino:       attr.Ino,
			Size:      attr.Size,
			Blocks:    attr.Blocks,
			Atime:     attr.Atime,
			Mtime:     attr.Mtime,
			Ctime:     attr.Ctime,
			Atimensec: attr.Atimensec,
			Mtimensec: attr.Mtimensec,
			Ctimensec: attr.Ctimensec,
			Mode:      attr.Mode,
			Nlink:     attr.Nlink,
			Owner:     &proto.Owner{Uid: attr.Owner.Uid, Gid: attr.Owner.Gid},
			Rdev:      attr.Rdev,
			Blksize:   attr.Blksize,
			Padding:   attr.Padding,
		},
		Status: int32(status),
	}
	return reply, nil
}

func (r *RpcServerImpl) OpenDir(ctx context.Context, request *proto.OpenDirRequest) (*proto.OpenDirReply, error) {
	dirs, s := r.filesystem.OpenDir(request.Path, createContext(ctx))
	// convert to proto.DirEntry
	entries := make([]*proto.DirEntry, len(dirs))
	for i, dir := range dirs {
		entries[i] = &proto.DirEntry{
			Name: dir.Name,
			Ino:  dir.Ino,
			Mode: dir.Mode,
			Off:  dir.Off,
		}
	}
	reply := &proto.OpenDirReply{
		Entries: entries,
		Status:  int32(s),
	}
	return reply, nil
}

func (r *RpcServerImpl) StatFs(ctx context.Context, request *proto.StatFsRequest) (*proto.StatFsReply, error) {
	statfs := r.filesystem.StatFs(request.Path)
	reply := &proto.StatFsReply{
		Blocks:  statfs.Blocks,
		Bfree:   statfs.Bfree,
		Bavail:  statfs.Bavail,
		Files:   statfs.Files,
		Ffree:   statfs.Ffree,
		Bsize:   statfs.Bsize,
		Namelen: statfs.NameLen,
		Frsize:  statfs.Frsize,
	}
	return reply, nil
}

// Unlink removes a file
func (r *RpcServerImpl) Unlink(ctx context.Context, request *proto.UnlinkRequest) (*proto.UnlinkReply, error) {
	status := r.filesystem.Unlink(request.Path, createContext(ctx))
	return &proto.UnlinkReply{Status: int32(status)}, nil
}

// createContext creates a new fuse.Context from the given context.Context
func createContext(ctx context.Context) *fuse.Context {
	return &fuse.Context{
		Caller: fuse.Caller{
			Owner: fuse.Owner{
				Uid: uint32(os.Getuid()),
				Gid: uint32(os.Getgid()),
			},
			Pid: uint32(os.Getpid()),
		},
		Cancel: ctx.Done(),
	}
}
