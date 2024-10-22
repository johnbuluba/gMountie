package io

import (
	"context"
	"gmountie/pkg/proto"
	"os"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"google.golang.org/grpc"
)

type RpcServerImpl struct {
	filesystem pathfs.FileSystem
	proto.UnimplementedRpcFsServer
}

// Verify that RpcServerImpl implements proto.RpcFsServer
var _ proto.RpcFsServer = (*RpcServerImpl)(nil)

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

// Mkdir creates a new directory
func (r *RpcServerImpl) Mkdir(ctx context.Context, request *proto.MkdirRequest) (*proto.MkdirReply, error) {
	status := r.filesystem.Mkdir(request.Path, request.Mode, createContext(ctx))
	return &proto.MkdirReply{Status: int32(status)}, nil
}

// Rmdir removes a directory
func (r *RpcServerImpl) Rmdir(ctx context.Context, request *proto.RmdirRequest) (*proto.RmdirReply, error) {
	status := r.filesystem.Rmdir(request.Path, createContext(ctx))
	return &proto.RmdirReply{Status: int32(status)}, nil
}

// OpenDir opens a directory
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

// Access checks if a file can be accessed
func (r *RpcServerImpl) Access(ctx context.Context, request *proto.AccessRequest) (*proto.AccessReply, error) {
	status := r.filesystem.Access(request.Path, request.Mode, createContext(ctx))
	return &proto.AccessReply{Status: int32(status)}, nil
}

// Truncate changes the size of a file
func (r *RpcServerImpl) Truncate(ctx context.Context, request *proto.TruncateRequest) (*proto.TruncateReply, error) {
	status := r.filesystem.Truncate(request.Path, request.Size, createContext(ctx))
	return &proto.TruncateReply{Status: int32(status)}, nil
}

// Chmod changes the mode of a file
func (r *RpcServerImpl) Chmod(ctx context.Context, request *proto.ChmodRequest) (*proto.ChmodReply, error) {
	status := r.filesystem.Chmod(request.Path, request.Mode, createContext(ctx))
	return &proto.ChmodReply{Status: int32(status)}, nil
}

// Chown changes the owner of a file
func (r *RpcServerImpl) Chown(ctx context.Context, request *proto.ChownRequest) (*proto.ChownReply, error) {
	status := r.filesystem.Chown(request.Path, request.Uid, request.Gid, createContext(ctx))
	return &proto.ChownReply{Status: int32(status)}, nil
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
