package io

import (
	"reflect"
	"runtime"

	"github.com/hanwen/go-fuse/v2/fuse/pathfs"
)

type Middleware func(pathfs.FileSystem) pathfs.FileSystem

// GetName returns the name of the middleware.
func (m Middleware) GetName() string {
	return runtime.FuncForPC(reflect.ValueOf(m).Pointer()).Name()
}
