// Code generated by mockery v2.46.3. DO NOT EDIT.

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	proto "gmountie/pkg/proto"
)

// MockRpcFileClient is an autogenerated mock type for the RpcFileClient type
type MockRpcFileClient struct {
	mock.Mock
}

type MockRpcFileClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRpcFileClient) EXPECT() *MockRpcFileClient_Expecter {
	return &MockRpcFileClient_Expecter{mock: &_m.Mock}
}

// Allocate provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Allocate(ctx context.Context, in *proto.AllocateRequest, opts ...grpc.CallOption) (*proto.AllocateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Allocate")
	}

	var r0 *proto.AllocateReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AllocateRequest, ...grpc.CallOption) (*proto.AllocateReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AllocateRequest, ...grpc.CallOption) *proto.AllocateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AllocateReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.AllocateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Allocate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Allocate'
type MockRpcFileClient_Allocate_Call struct {
	*mock.Call
}

// Allocate is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.AllocateRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Allocate(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Allocate_Call {
	return &MockRpcFileClient_Allocate_Call{Call: _e.mock.On("Allocate",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Allocate_Call) Run(run func(ctx context.Context, in *proto.AllocateRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Allocate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.AllocateRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Allocate_Call) Return(_a0 *proto.AllocateReply, _a1 error) *MockRpcFileClient_Allocate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Allocate_Call) RunAndReturn(run func(context.Context, *proto.AllocateRequest, ...grpc.CallOption) (*proto.AllocateReply, error)) *MockRpcFileClient_Allocate_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Create(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption) (*proto.CreateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *proto.CreateReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateRequest, ...grpc.CallOption) (*proto.CreateReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.CreateRequest, ...grpc.CallOption) *proto.CreateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.CreateReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRpcFileClient_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.CreateRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Create(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Create_Call {
	return &MockRpcFileClient_Create_Call{Call: _e.mock.On("Create",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Create_Call) Run(run func(ctx context.Context, in *proto.CreateRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.CreateRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Create_Call) Return(_a0 *proto.CreateReply, _a1 error) *MockRpcFileClient_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Create_Call) RunAndReturn(run func(context.Context, *proto.CreateRequest, ...grpc.CallOption) (*proto.CreateReply, error)) *MockRpcFileClient_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Flush provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Flush(ctx context.Context, in *proto.FlushRequest, opts ...grpc.CallOption) (*proto.FlushReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Flush")
	}

	var r0 *proto.FlushReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.FlushRequest, ...grpc.CallOption) (*proto.FlushReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.FlushRequest, ...grpc.CallOption) *proto.FlushReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.FlushReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.FlushRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Flush_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Flush'
type MockRpcFileClient_Flush_Call struct {
	*mock.Call
}

// Flush is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.FlushRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Flush(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Flush_Call {
	return &MockRpcFileClient_Flush_Call{Call: _e.mock.On("Flush",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Flush_Call) Run(run func(ctx context.Context, in *proto.FlushRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Flush_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.FlushRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Flush_Call) Return(_a0 *proto.FlushReply, _a1 error) *MockRpcFileClient_Flush_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Flush_Call) RunAndReturn(run func(context.Context, *proto.FlushRequest, ...grpc.CallOption) (*proto.FlushReply, error)) *MockRpcFileClient_Flush_Call {
	_c.Call.Return(run)
	return _c
}

// Fsync provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Fsync(ctx context.Context, in *proto.FsyncRequest, opts ...grpc.CallOption) (*proto.FsyncReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Fsync")
	}

	var r0 *proto.FsyncReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.FsyncRequest, ...grpc.CallOption) (*proto.FsyncReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.FsyncRequest, ...grpc.CallOption) *proto.FsyncReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.FsyncReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.FsyncRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Fsync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fsync'
type MockRpcFileClient_Fsync_Call struct {
	*mock.Call
}

// Fsync is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.FsyncRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Fsync(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Fsync_Call {
	return &MockRpcFileClient_Fsync_Call{Call: _e.mock.On("Fsync",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Fsync_Call) Run(run func(ctx context.Context, in *proto.FsyncRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Fsync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.FsyncRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Fsync_Call) Return(_a0 *proto.FsyncReply, _a1 error) *MockRpcFileClient_Fsync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Fsync_Call) RunAndReturn(run func(context.Context, *proto.FsyncRequest, ...grpc.CallOption) (*proto.FsyncReply, error)) *MockRpcFileClient_Fsync_Call {
	_c.Call.Return(run)
	return _c
}

// GetLk provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) GetLk(ctx context.Context, in *proto.GetLkRequest, opts ...grpc.CallOption) (*proto.GetLkReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetLk")
	}

	var r0 *proto.GetLkReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetLkRequest, ...grpc.CallOption) (*proto.GetLkReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetLkRequest, ...grpc.CallOption) *proto.GetLkReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetLkReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetLkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_GetLk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLk'
type MockRpcFileClient_GetLk_Call struct {
	*mock.Call
}

// GetLk is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.GetLkRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) GetLk(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_GetLk_Call {
	return &MockRpcFileClient_GetLk_Call{Call: _e.mock.On("GetLk",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_GetLk_Call) Run(run func(ctx context.Context, in *proto.GetLkRequest, opts ...grpc.CallOption)) *MockRpcFileClient_GetLk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.GetLkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_GetLk_Call) Return(_a0 *proto.GetLkReply, _a1 error) *MockRpcFileClient_GetLk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_GetLk_Call) RunAndReturn(run func(context.Context, *proto.GetLkRequest, ...grpc.CallOption) (*proto.GetLkReply, error)) *MockRpcFileClient_GetLk_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Open(ctx context.Context, in *proto.OpenRequest, opts ...grpc.CallOption) (*proto.OpenReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Open")
	}

	var r0 *proto.OpenReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenRequest, ...grpc.CallOption) (*proto.OpenReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenRequest, ...grpc.CallOption) *proto.OpenReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.OpenReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.OpenRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MockRpcFileClient_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.OpenRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Open(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Open_Call {
	return &MockRpcFileClient_Open_Call{Call: _e.mock.On("Open",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Open_Call) Run(run func(ctx context.Context, in *proto.OpenRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.OpenRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Open_Call) Return(_a0 *proto.OpenReply, _a1 error) *MockRpcFileClient_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Open_Call) RunAndReturn(run func(context.Context, *proto.OpenRequest, ...grpc.CallOption) (*proto.OpenReply, error)) *MockRpcFileClient_Open_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Read(ctx context.Context, in *proto.ReadRequest, opts ...grpc.CallOption) (*proto.ReadReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 *proto.ReadReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ReadRequest, ...grpc.CallOption) (*proto.ReadReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ReadRequest, ...grpc.CallOption) *proto.ReadReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ReadReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ReadRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockRpcFileClient_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.ReadRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Read(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Read_Call {
	return &MockRpcFileClient_Read_Call{Call: _e.mock.On("Read",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Read_Call) Run(run func(ctx context.Context, in *proto.ReadRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.ReadRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Read_Call) Return(_a0 *proto.ReadReply, _a1 error) *MockRpcFileClient_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Read_Call) RunAndReturn(run func(context.Context, *proto.ReadRequest, ...grpc.CallOption) (*proto.ReadReply, error)) *MockRpcFileClient_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Release provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Release(ctx context.Context, in *proto.ReleaseRequest, opts ...grpc.CallOption) (*proto.ReleaseReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Release")
	}

	var r0 *proto.ReleaseReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ReleaseRequest, ...grpc.CallOption) (*proto.ReleaseReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ReleaseRequest, ...grpc.CallOption) *proto.ReleaseReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ReleaseReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ReleaseRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Release_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Release'
type MockRpcFileClient_Release_Call struct {
	*mock.Call
}

// Release is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.ReleaseRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Release(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Release_Call {
	return &MockRpcFileClient_Release_Call{Call: _e.mock.On("Release",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Release_Call) Run(run func(ctx context.Context, in *proto.ReleaseRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Release_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.ReleaseRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Release_Call) Return(_a0 *proto.ReleaseReply, _a1 error) *MockRpcFileClient_Release_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Release_Call) RunAndReturn(run func(context.Context, *proto.ReleaseRequest, ...grpc.CallOption) (*proto.ReleaseReply, error)) *MockRpcFileClient_Release_Call {
	_c.Call.Return(run)
	return _c
}

// SetLk provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) SetLk(ctx context.Context, in *proto.SetLkRequest, opts ...grpc.CallOption) (*proto.SetLkReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetLk")
	}

	var r0 *proto.SetLkReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SetLkRequest, ...grpc.CallOption) (*proto.SetLkReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SetLkRequest, ...grpc.CallOption) *proto.SetLkReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SetLkReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.SetLkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_SetLk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLk'
type MockRpcFileClient_SetLk_Call struct {
	*mock.Call
}

// SetLk is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.SetLkRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) SetLk(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_SetLk_Call {
	return &MockRpcFileClient_SetLk_Call{Call: _e.mock.On("SetLk",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_SetLk_Call) Run(run func(ctx context.Context, in *proto.SetLkRequest, opts ...grpc.CallOption)) *MockRpcFileClient_SetLk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.SetLkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_SetLk_Call) Return(_a0 *proto.SetLkReply, _a1 error) *MockRpcFileClient_SetLk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_SetLk_Call) RunAndReturn(run func(context.Context, *proto.SetLkRequest, ...grpc.CallOption) (*proto.SetLkReply, error)) *MockRpcFileClient_SetLk_Call {
	_c.Call.Return(run)
	return _c
}

// SetLkw provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) SetLkw(ctx context.Context, in *proto.SetLkwRequest, opts ...grpc.CallOption) (*proto.SetLkwReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SetLkw")
	}

	var r0 *proto.SetLkwReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SetLkwRequest, ...grpc.CallOption) (*proto.SetLkwReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SetLkwRequest, ...grpc.CallOption) *proto.SetLkwReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SetLkwReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.SetLkwRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_SetLkw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetLkw'
type MockRpcFileClient_SetLkw_Call struct {
	*mock.Call
}

// SetLkw is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.SetLkwRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) SetLkw(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_SetLkw_Call {
	return &MockRpcFileClient_SetLkw_Call{Call: _e.mock.On("SetLkw",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_SetLkw_Call) Run(run func(ctx context.Context, in *proto.SetLkwRequest, opts ...grpc.CallOption)) *MockRpcFileClient_SetLkw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.SetLkwRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_SetLkw_Call) Return(_a0 *proto.SetLkwReply, _a1 error) *MockRpcFileClient_SetLkw_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_SetLkw_Call) RunAndReturn(run func(context.Context, *proto.SetLkwRequest, ...grpc.CallOption) (*proto.SetLkwReply, error)) *MockRpcFileClient_SetLkw_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFileClient) Write(ctx context.Context, in *proto.WriteRequest, opts ...grpc.CallOption) (*proto.WriteReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 *proto.WriteReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WriteRequest, ...grpc.CallOption) (*proto.WriteReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WriteRequest, ...grpc.CallOption) *proto.WriteReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.WriteReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.WriteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFileClient_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockRpcFileClient_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.WriteRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFileClient_Expecter) Write(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFileClient_Write_Call {
	return &MockRpcFileClient_Write_Call{Call: _e.mock.On("Write",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFileClient_Write_Call) Run(run func(ctx context.Context, in *proto.WriteRequest, opts ...grpc.CallOption)) *MockRpcFileClient_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.WriteRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFileClient_Write_Call) Return(_a0 *proto.WriteReply, _a1 error) *MockRpcFileClient_Write_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFileClient_Write_Call) RunAndReturn(run func(context.Context, *proto.WriteRequest, ...grpc.CallOption) (*proto.WriteReply, error)) *MockRpcFileClient_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRpcFileClient creates a new instance of MockRpcFileClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRpcFileClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRpcFileClient {
	mock := &MockRpcFileClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}