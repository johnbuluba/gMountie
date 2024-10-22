// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/proto/fs.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcFsClient is the client API for RpcFs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcFsClient interface {
	GetAttr(ctx context.Context, in *GetAttrRequest, opts ...grpc.CallOption) (*GetAttrReply, error)
	StatFs(ctx context.Context, in *StatFsRequest, opts ...grpc.CallOption) (*StatFsReply, error)
	OpenDir(ctx context.Context, in *OpenDirRequest, opts ...grpc.CallOption) (*OpenDirReply, error)
	Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*UnlinkReply, error)
	Access(ctx context.Context, in *AccessRequest, opts ...grpc.CallOption) (*AccessReply, error)
	Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateReply, error)
	Chown(ctx context.Context, in *ChownRequest, opts ...grpc.CallOption) (*ChownReply, error)
	Chmod(ctx context.Context, in *ChmodRequest, opts ...grpc.CallOption) (*ChmodReply, error)
	Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirReply, error)
	Rmdir(ctx context.Context, in *RmdirRequest, opts ...grpc.CallOption) (*RmdirReply, error)
	Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*RenameReply, error)
}

type rpcFsClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcFsClient(cc grpc.ClientConnInterface) RpcFsClient {
	return &rpcFsClient{cc}
}

func (c *rpcFsClient) GetAttr(ctx context.Context, in *GetAttrRequest, opts ...grpc.CallOption) (*GetAttrReply, error) {
	out := new(GetAttrReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/GetAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) StatFs(ctx context.Context, in *StatFsRequest, opts ...grpc.CallOption) (*StatFsReply, error) {
	out := new(StatFsReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/StatFs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) OpenDir(ctx context.Context, in *OpenDirRequest, opts ...grpc.CallOption) (*OpenDirReply, error) {
	out := new(OpenDirReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/OpenDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*UnlinkReply, error) {
	out := new(UnlinkReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Unlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Access(ctx context.Context, in *AccessRequest, opts ...grpc.CallOption) (*AccessReply, error) {
	out := new(AccessReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Access", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*TruncateReply, error) {
	out := new(TruncateReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Truncate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Chown(ctx context.Context, in *ChownRequest, opts ...grpc.CallOption) (*ChownReply, error) {
	out := new(ChownReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Chown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Chmod(ctx context.Context, in *ChmodRequest, opts ...grpc.CallOption) (*ChmodReply, error) {
	out := new(ChmodReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Chmod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirReply, error) {
	out := new(MkdirReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Rmdir(ctx context.Context, in *RmdirRequest, opts ...grpc.CallOption) (*RmdirReply, error) {
	out := new(RmdirReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Rmdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcFsClient) Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*RenameReply, error) {
	out := new(RenameReply)
	err := c.cc.Invoke(ctx, "/gmountie.RpcFs/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcFsServer is the server API for RpcFs service.
// All implementations must embed UnimplementedRpcFsServer
// for forward compatibility
type RpcFsServer interface {
	GetAttr(context.Context, *GetAttrRequest) (*GetAttrReply, error)
	StatFs(context.Context, *StatFsRequest) (*StatFsReply, error)
	OpenDir(context.Context, *OpenDirRequest) (*OpenDirReply, error)
	Unlink(context.Context, *UnlinkRequest) (*UnlinkReply, error)
	Access(context.Context, *AccessRequest) (*AccessReply, error)
	Truncate(context.Context, *TruncateRequest) (*TruncateReply, error)
	Chown(context.Context, *ChownRequest) (*ChownReply, error)
	Chmod(context.Context, *ChmodRequest) (*ChmodReply, error)
	Mkdir(context.Context, *MkdirRequest) (*MkdirReply, error)
	Rmdir(context.Context, *RmdirRequest) (*RmdirReply, error)
	Rename(context.Context, *RenameRequest) (*RenameReply, error)
	mustEmbedUnimplementedRpcFsServer()
}

// UnimplementedRpcFsServer must be embedded to have forward compatible implementations.
type UnimplementedRpcFsServer struct {
}

func (UnimplementedRpcFsServer) GetAttr(context.Context, *GetAttrRequest) (*GetAttrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttr not implemented")
}
func (UnimplementedRpcFsServer) StatFs(context.Context, *StatFsRequest) (*StatFsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatFs not implemented")
}
func (UnimplementedRpcFsServer) OpenDir(context.Context, *OpenDirRequest) (*OpenDirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenDir not implemented")
}
func (UnimplementedRpcFsServer) Unlink(context.Context, *UnlinkRequest) (*UnlinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlink not implemented")
}
func (UnimplementedRpcFsServer) Access(context.Context, *AccessRequest) (*AccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Access not implemented")
}
func (UnimplementedRpcFsServer) Truncate(context.Context, *TruncateRequest) (*TruncateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Truncate not implemented")
}
func (UnimplementedRpcFsServer) Chown(context.Context, *ChownRequest) (*ChownReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chown not implemented")
}
func (UnimplementedRpcFsServer) Chmod(context.Context, *ChmodRequest) (*ChmodReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chmod not implemented")
}
func (UnimplementedRpcFsServer) Mkdir(context.Context, *MkdirRequest) (*MkdirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (UnimplementedRpcFsServer) Rmdir(context.Context, *RmdirRequest) (*RmdirReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rmdir not implemented")
}
func (UnimplementedRpcFsServer) Rename(context.Context, *RenameRequest) (*RenameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (UnimplementedRpcFsServer) mustEmbedUnimplementedRpcFsServer() {}

// UnsafeRpcFsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcFsServer will
// result in compilation errors.
type UnsafeRpcFsServer interface {
	mustEmbedUnimplementedRpcFsServer()
}

func RegisterRpcFsServer(s grpc.ServiceRegistrar, srv RpcFsServer) {
	s.RegisterService(&RpcFs_ServiceDesc, srv)
}

func _RpcFs_GetAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).GetAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/GetAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).GetAttr(ctx, req.(*GetAttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_StatFs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatFsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).StatFs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/StatFs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).StatFs(ctx, req.(*StatFsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_OpenDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).OpenDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/OpenDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).OpenDir(ctx, req.(*OpenDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Unlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Unlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Unlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Unlink(ctx, req.(*UnlinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Access_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Access(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Access",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Access(ctx, req.(*AccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruncateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Truncate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Truncate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Truncate(ctx, req.(*TruncateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Chown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Chown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Chown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Chown(ctx, req.(*ChownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Chmod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChmodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Chmod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Chmod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Chmod(ctx, req.(*ChmodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Mkdir(ctx, req.(*MkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Rmdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Rmdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Rmdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Rmdir(ctx, req.(*RmdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcFs_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcFsServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmountie.RpcFs/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcFsServer).Rename(ctx, req.(*RenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcFs_ServiceDesc is the grpc.ServiceDesc for RpcFs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcFs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gmountie.RpcFs",
	HandlerType: (*RpcFsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAttr",
			Handler:    _RpcFs_GetAttr_Handler,
		},
		{
			MethodName: "StatFs",
			Handler:    _RpcFs_StatFs_Handler,
		},
		{
			MethodName: "OpenDir",
			Handler:    _RpcFs_OpenDir_Handler,
		},
		{
			MethodName: "Unlink",
			Handler:    _RpcFs_Unlink_Handler,
		},
		{
			MethodName: "Access",
			Handler:    _RpcFs_Access_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _RpcFs_Truncate_Handler,
		},
		{
			MethodName: "Chown",
			Handler:    _RpcFs_Chown_Handler,
		},
		{
			MethodName: "Chmod",
			Handler:    _RpcFs_Chmod_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _RpcFs_Mkdir_Handler,
		},
		{
			MethodName: "Rmdir",
			Handler:    _RpcFs_Rmdir_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _RpcFs_Rename_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/fs.proto",
}
