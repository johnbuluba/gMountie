package io

import (
	"context"
	"grpc-fs/pkg/common"
	"grpc-fs/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
	"go.uber.org/zap"
)

type GrpcInode struct {
	fsClient   proto.RpcFsClient
	fileClient proto.RpcFileClient
	pathfs.FileSystem
}

// NewGrpcInode creates a new GrpcInode
func NewGrpcInode(fsClient proto.RpcFsClient, fileClient proto.RpcFileClient) pathfs.FileSystem {
	return &GrpcInode{
		fsClient:   fsClient,
		fileClient: fileClient,
		FileSystem: pathfs.NewDefaultFileSystem(),
	}
}

func (fs *GrpcInode) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	res, err := fs.fsClient.GetAttr(context, &proto.GetAttrRequest{Path: name})
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
func (fs *GrpcInode) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, code fuse.Status) {
	res, err := fs.fsClient.OpenDir(context, &proto.OpenDirRequest{Path: name})
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
	res, err := fs.fileClient.Open(context, &proto.OpenRequest{Path: name, Flags: flags})
	if err != nil || res == nil {
		common.Log.Error("error in call: Open", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	if fuse.Status(res.Status) != fuse.OK {
		return nil, fuse.Status(res.Status)
	}
	return NewGrpcFile(fs.fileClient, name), fuse.Status(res.Status)
}

func (fs *GrpcInode) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	res, err := fs.fileClient.Create(context, &proto.CreateRequest{Path: name, Flags: flags, Mode: mode})
	if err != nil || res == nil {
		common.Log.Error("error in call: Create", zap.String("path", name), zap.Error(err))
		return nil, fuse.EIO
	}
	if fuse.Status(res.Status) != fuse.OK {
		return nil, fuse.Status(res.Status)
	}
	return NewGrpcFile(fs.fileClient, name), fuse.Status(res.Status)
}
func (fs *GrpcInode) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	res, err := fs.fsClient.Unlink(context, &proto.UnlinkRequest{Path: name})
	if err != nil || res == nil {
		common.Log.Error("error in call: Unlink", zap.String("path", name), zap.Error(err))
		return fuse.EIO
	}
	return fuse.Status(res.Status)
}

func (fs *GrpcInode) StatFs(name string) *fuse.StatfsOut {
	res, err := fs.fsClient.StatFs(context.Background(), &proto.StatFsRequest{Path: name})
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
