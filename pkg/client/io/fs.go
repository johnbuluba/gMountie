package io

import (
	"context"
	"gmountie/pkg/client/grpc"
	"gmountie/pkg/common"
	"gmountie/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"go.uber.org/zap"
)

type GrpcInode struct {
	client *grpc.Client
	pathfs.FileSystem
}

// NewGrpcInode creates a new GrpcInode
func NewGrpcInode(client *grpc.Client) pathfs.FileSystem {
	return &GrpcInode{
		client:     client,
		FileSystem: pathfs.NewDefaultFileSystem(),
	}
}

// OnMount is called after the file system is mounted
func (fs *GrpcInode) OnMount(nodeFs *pathfs.PathNodeFs) {
	common.Log.Info("File system is mounted")
}

// GetAttr returns the attributes of a file
func (fs *GrpcInode) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	res, err := fs.client.Fs.GetAttr(context, &proto.GetAttrRequest{Caller: createCaller(context), Path: name})
	if err != nil {
		common.Log.Error("error in call: GetAttr", zap.String("path", name), zap.Error(err))
		return &fuse.Attr{}, fuse.EIO
	}
	if res.GetAttributes() == nil {
		return &fuse.Attr{}, fuse.Status(res.Status)
	}
	a := &fuse.Attr{
		Ino:    res.GetAttributes().Ino,
		Size:   res.GetAttributes().Size,
		Blocks: res.GetAttributes().Blocks,
		Atime:  res.GetAttributes().Atime,
		Mtime:  res.GetAttributes().Mtime,
		Ctime:  res.GetAttributes().Ctime,
		Mode:   res.GetAttributes().Mode,
		Nlink:  res.GetAttributes().Nlink,
		Owner: fuse.Owner{
			Uid: res.GetAttributes().Owner.Uid,
			Gid: res.GetAttributes().Owner.Gid,
		},
		Rdev:    res.GetAttributes().Rdev,
		Blksize: res.GetAttributes().Blksize,
		Padding: res.GetAttributes().Padding,
	}
	return a, fuse.Status(res.Status)
}

// Mkdir creates a directory
func (fs *GrpcInode) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	res, err := fs.client.Fs.Mkdir(context, &proto.MkdirRequest{Path: name, Mode: mode})
	if err != nil || res == nil {
		common.Log.Error("error in call: MkDir", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// Rmdir removes a directory
func (fs *GrpcInode) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Rmdir(context, &proto.RmdirRequest{Caller: createCaller(context), Path: name})
	if err != nil || res == nil {
		common.Log.Error("error in call: RmDir", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// Rename renames a file
func (fs *GrpcInode) Rename(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Rename(context, &proto.RenameRequest{Caller: createCaller(context), OldName: oldName, NewName: newName})
	if err != nil || res == nil {
		common.Log.Error("error in call: Rename", zap.String("oldName", oldName), zap.String("newName", newName), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// OpenDir opens a directory
func (fs *GrpcInode) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, code fuse.Status) {
	res, err := fs.client.Fs.OpenDir(context, &proto.OpenDirRequest{Caller: createCaller(context), Path: name})
	if err != nil || res == nil {
		common.Log.Error("error in call: OpenDir", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	var entries []fuse.DirEntry
	for _, entry := range res.Entries {
		entries = append(entries, fuse.DirEntry{
			Mode: entry.Mode,
			Name: entry.Name,
			Ino:  entry.Ino,
			Off:  entry.Off,
		})
	}
	return entries, fuse.Status(res.Status)
}

func (fs *GrpcInode) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	res, err := fs.client.File.Open(context, &proto.OpenRequest{Caller: createCaller(context), Path: name, Flags: flags})
	if err != nil || res == nil {
		common.Log.Error("error in call: Open", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	if fuse.Status(res.Status) != fuse.OK {
		return nil, fuse.Status(res.Status)
	}
	return NewGrpcFile(fs.client.File, name), fuse.Status(res.Status)
}

func (fs *GrpcInode) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	res, err := fs.client.File.Create(context, &proto.CreateRequest{Caller: createCaller(context), Path: name, Flags: flags, Mode: mode})
	if err != nil || res == nil {
		common.Log.Error("error in call: Create", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	if fuse.Status(res.Status) != fuse.OK {
		return nil, fuse.Status(res.Status)
	}
	return NewGrpcFile(fs.client.File, name), fuse.Status(res.Status)
}

func (fs *GrpcInode) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Unlink(context, &proto.UnlinkRequest{Caller: createCaller(context), Path: name})
	if err != nil || res == nil {
		common.Log.Error("error in call: Unlink", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

func (fs *GrpcInode) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Access(context, &proto.AccessRequest{Caller: createCaller(context), Path: name, Mode: mode})
	if err != nil || res == nil {
		common.Log.Error("error in call: Access", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// Truncate truncates a file
func (fs *GrpcInode) Truncate(name string, size uint64, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Truncate(context, &proto.TruncateRequest{Caller: createCaller(context), Path: name, Size: size})
	if err != nil || res == nil {
		common.Log.Error("error in call: Truncate", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// Chmod changes the mode of a file
func (fs *GrpcInode) Chmod(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Chmod(context, &proto.ChmodRequest{Caller: createCaller(context), Path: name, Mode: mode})
	if err != nil || res == nil {
		common.Log.Error("error in call: Chmod", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// Chown changes the owner of a file
func (fs *GrpcInode) Chown(name string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	res, err := fs.client.Fs.Chown(context, &proto.ChownRequest{Caller: createCaller(context), Path: name, Uid: uid, Gid: gid})
	if err != nil || res == nil {
		common.Log.Error("error in call: Chown", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

// ------------------- Extended attributes -------------------

// GetXAttr gets an extended attribute
func (fs *GrpcInode) GetXAttr(name string, attribute string, context *fuse.Context) (data []byte, code fuse.Status) {
	res, err := fs.client.Fs.GetXAttr(context, &proto.GetXAttrRequest{Caller: createCaller(context), Path: name, Attribute: attribute})
	if err != nil || res == nil {
		common.Log.Error("error in call: GetXAttr", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	return res.Data, fuse.Status(res.Status)
}

func (fs *GrpcInode) StatFs(name string) *fuse.StatfsOut {
	res, err := fs.client.Fs.StatFs(context.Background(), &proto.StatFsRequest{Path: name})
	if err != nil || res == nil {
		common.Log.Error("error in call: StatFs", zap.String("path", name), zap.Error(err))
		return nil
	}
	stats := &fuse.StatfsOut{
		Blocks:  res.Blocks,
		Bfree:   res.Bfree,
		Bavail:  res.Bavail,
		Files:   res.Files,
		Ffree:   res.Ffree,
		Bsize:   res.Bsize,
		NameLen: res.Namelen,
		Frsize:  res.Frsize,
	}
	if len(res.Spare) == 6 {
		stats.Spare = [6]uint32{res.Spare[0], res.Spare[1], res.Spare[2], res.Spare[3], res.Spare[4], res.Spare[5]}
	}
	return stats
}

// createCaller creates a caller for the file system
func createCaller(ctx *fuse.Context) *proto.Caller {
	return &proto.Caller{
		Owner: &proto.Owner{
			Uid: ctx.Owner.Uid,
			Gid: ctx.Owner.Gid,
		},
		Pid: ctx.Pid,
	}
}
