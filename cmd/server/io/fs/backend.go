package fs

import (
	"github.com/hanwen/go-fuse/v2/fs"
)

// Backend is implementing a backend for the filesystem.
type Backend struct {
	fs.InodeEmbedder
}

// NewBackend creates a new filesystem backend.
func NewBackend(root string) (*Backend, error) {
	embedder, err := fs.NewLoopbackRoot(root)
	if err != nil {
		return nil, err

	}
	embedder.EmbeddedInode()
	return &Backend{
		InodeEmbedder: embedder,
	}, nil
}
