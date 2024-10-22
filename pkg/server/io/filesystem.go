package io

import "github.com/hanwen/go-fuse/v2/fuse/pathfs"

// LocalFilesystem is a struct that contains information about a local filesystem.
type LocalFilesystem struct {
	pathfs.FileSystem
}

// NewLocalFilesystem creates a new LocalFilesystem.
func NewLocalFilesystem(path string) *LocalFilesystem {
	return &LocalFilesystem{
		FileSystem: pathfs.NewLoopbackFileSystem(path),
	}
}
