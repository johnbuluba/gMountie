package controller

import (
	"context"
	"gmountie/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
)

// createContext creates a new fuse.Context from the given context.Context
func createContext(ctx context.Context, caller *proto.Caller) *fuse.Context {
	return &fuse.Context{
		Caller: fuse.Caller{
			Owner: fuse.Owner{
				Uid: caller.Owner.Uid,
				Gid: caller.Owner.Gid,
			},
			Pid: caller.Pid,
		},
		Cancel: ctx.Done(),
	}
}
