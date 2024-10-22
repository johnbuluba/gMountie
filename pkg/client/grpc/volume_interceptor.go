package grpc

import (
	"context"
	"gmountie/pkg/common"

	"google.golang.org/grpc"
)

// NewVolumeInterceptor creates a new volume interceptor.
func NewVolumeInterceptor(volume string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = common.AddVolumeHeader(ctx, volume)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
