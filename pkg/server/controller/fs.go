package controller

import (
	"context"
	"gmountie/pkg/proto"
	"gmountie/pkg/server/service"

	"google.golang.org/grpc"
)

type RpcServerImpl struct {
	fsService service.VolumeService
	proto.UnimplementedRpcFsServer
}

// Verify that RpcServerImpl implements proto.RpcFsServer
var _ proto.RpcFsServer = (*RpcServerImpl)(nil)

// NewGrpcServer creates a new gRPC server
func NewGrpcServer(fsService service.VolumeService) *RpcServerImpl {
	return &RpcServerImpl{
		fsService: fsService,
	}
}

// Register registers the gRPC server
func (r *RpcServerImpl) Register(server *grpc.Server) {
	proto.RegisterRpcFsServer(server, r)
}

func (r *RpcServerImpl) GetAttr(ctx context.Context, request *proto.GetAttrRequest) (*proto.GetAttrReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	attr, status := fs.GetAttr(request.Path, createContext(ctx, request.Caller))
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
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Mkdir(request.Path, request.Mode, createContext(ctx, request.Caller))
	return &proto.MkdirReply{Status: int32(status)}, nil
}

// Rmdir removes a directory
func (r *RpcServerImpl) Rmdir(ctx context.Context, request *proto.RmdirRequest) (*proto.RmdirReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Rmdir(request.Path, createContext(ctx, request.Caller))
	return &proto.RmdirReply{Status: int32(status)}, nil
}

// Rename renames a file
func (r *RpcServerImpl) Rename(ctx context.Context, request *proto.RenameRequest) (*proto.RenameReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Rename(request.OldName, request.NewName, createContext(ctx, request.Caller))
	return &proto.RenameReply{Status: int32(status)}, nil
}

// OpenDir opens a directory
func (r *RpcServerImpl) OpenDir(ctx context.Context, request *proto.OpenDirRequest) (*proto.OpenDirReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	dirs, s := fs.OpenDir(request.Path, createContext(ctx, request.Caller))
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
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	statfs := fs.StatFs(request.Path)
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
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Unlink(request.Path, createContext(ctx, request.Caller))
	return &proto.UnlinkReply{Status: int32(status)}, nil
}

// Access checks if a file can be accessed
func (r *RpcServerImpl) Access(ctx context.Context, request *proto.AccessRequest) (*proto.AccessReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Access(request.Path, request.Mode, createContext(ctx, request.Caller))
	return &proto.AccessReply{Status: int32(status)}, nil
}

// Truncate changes the size of a file
func (r *RpcServerImpl) Truncate(ctx context.Context, request *proto.TruncateRequest) (*proto.TruncateReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Truncate(request.Path, request.Size, createContext(ctx, request.Caller))
	return &proto.TruncateReply{Status: int32(status)}, nil
}

// Chmod changes the mode of a file
func (r *RpcServerImpl) Chmod(ctx context.Context, request *proto.ChmodRequest) (*proto.ChmodReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Chmod(request.Path, request.Mode, createContext(ctx, request.Caller))
	return &proto.ChmodReply{Status: int32(status)}, nil
}

// Chown changes the owner of a file
func (r *RpcServerImpl) Chown(ctx context.Context, request *proto.ChownRequest) (*proto.ChownReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	status := fs.Chown(request.Path, request.Uid, request.Gid, createContext(ctx, request.Caller))
	return &proto.ChownReply{Status: int32(status)}, nil
}

// ----- Extended attributes -----

// GetXAttr gets an extended attribute
func (r *RpcServerImpl) GetXAttr(ctx context.Context, request *proto.GetXAttrRequest) (*proto.GetXAttrReply, error) {
	fs, err := r.fsService.GetVolumeFileSystemFromContext(ctx)
	if err != nil {
		return nil, err
	}
	data, status := fs.GetXAttr(request.Path, request.Attribute, createContext(ctx, request.Caller))
	return &proto.GetXAttrReply{Data: data, Status: int32(status)}, nil
}
