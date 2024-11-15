// Code generated by mockery v2.46.3. DO NOT EDIT.

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	proto "gmountie/pkg/proto"
)

// MockRpcFsClient is an autogenerated mock type for the RpcFsClient type
type MockRpcFsClient struct {
	mock.Mock
}

type MockRpcFsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRpcFsClient) EXPECT() *MockRpcFsClient_Expecter {
	return &MockRpcFsClient_Expecter{mock: &_m.Mock}
}

// Access provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Access(ctx context.Context, in *proto.AccessRequest, opts ...grpc.CallOption) (*proto.AccessReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Access")
	}

	var r0 *proto.AccessReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AccessRequest, ...grpc.CallOption) (*proto.AccessReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AccessRequest, ...grpc.CallOption) *proto.AccessReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AccessReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.AccessRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Access_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Access'
type MockRpcFsClient_Access_Call struct {
	*mock.Call
}

// Access is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.AccessRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Access(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Access_Call {
	return &MockRpcFsClient_Access_Call{Call: _e.mock.On("Access",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Access_Call) Run(run func(ctx context.Context, in *proto.AccessRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Access_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.AccessRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Access_Call) Return(_a0 *proto.AccessReply, _a1 error) *MockRpcFsClient_Access_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Access_Call) RunAndReturn(run func(context.Context, *proto.AccessRequest, ...grpc.CallOption) (*proto.AccessReply, error)) *MockRpcFsClient_Access_Call {
	_c.Call.Return(run)
	return _c
}

// Chmod provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Chmod(ctx context.Context, in *proto.ChmodRequest, opts ...grpc.CallOption) (*proto.ChmodReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Chmod")
	}

	var r0 *proto.ChmodReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChmodRequest, ...grpc.CallOption) (*proto.ChmodReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChmodRequest, ...grpc.CallOption) *proto.ChmodReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ChmodReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ChmodRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Chmod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chmod'
type MockRpcFsClient_Chmod_Call struct {
	*mock.Call
}

// Chmod is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.ChmodRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Chmod(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Chmod_Call {
	return &MockRpcFsClient_Chmod_Call{Call: _e.mock.On("Chmod",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Chmod_Call) Run(run func(ctx context.Context, in *proto.ChmodRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Chmod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.ChmodRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Chmod_Call) Return(_a0 *proto.ChmodReply, _a1 error) *MockRpcFsClient_Chmod_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Chmod_Call) RunAndReturn(run func(context.Context, *proto.ChmodRequest, ...grpc.CallOption) (*proto.ChmodReply, error)) *MockRpcFsClient_Chmod_Call {
	_c.Call.Return(run)
	return _c
}

// Chown provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Chown(ctx context.Context, in *proto.ChownRequest, opts ...grpc.CallOption) (*proto.ChownReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Chown")
	}

	var r0 *proto.ChownReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChownRequest, ...grpc.CallOption) (*proto.ChownReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChownRequest, ...grpc.CallOption) *proto.ChownReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ChownReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ChownRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Chown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chown'
type MockRpcFsClient_Chown_Call struct {
	*mock.Call
}

// Chown is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.ChownRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Chown(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Chown_Call {
	return &MockRpcFsClient_Chown_Call{Call: _e.mock.On("Chown",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Chown_Call) Run(run func(ctx context.Context, in *proto.ChownRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Chown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.ChownRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Chown_Call) Return(_a0 *proto.ChownReply, _a1 error) *MockRpcFsClient_Chown_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Chown_Call) RunAndReturn(run func(context.Context, *proto.ChownRequest, ...grpc.CallOption) (*proto.ChownReply, error)) *MockRpcFsClient_Chown_Call {
	_c.Call.Return(run)
	return _c
}

// GetAttr provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) GetAttr(ctx context.Context, in *proto.GetAttrRequest, opts ...grpc.CallOption) (*proto.GetAttrReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAttr")
	}

	var r0 *proto.GetAttrReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetAttrRequest, ...grpc.CallOption) (*proto.GetAttrReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetAttrRequest, ...grpc.CallOption) *proto.GetAttrReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetAttrReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetAttrRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_GetAttr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAttr'
type MockRpcFsClient_GetAttr_Call struct {
	*mock.Call
}

// GetAttr is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.GetAttrRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) GetAttr(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_GetAttr_Call {
	return &MockRpcFsClient_GetAttr_Call{Call: _e.mock.On("GetAttr",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_GetAttr_Call) Run(run func(ctx context.Context, in *proto.GetAttrRequest, opts ...grpc.CallOption)) *MockRpcFsClient_GetAttr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.GetAttrRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_GetAttr_Call) Return(_a0 *proto.GetAttrReply, _a1 error) *MockRpcFsClient_GetAttr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_GetAttr_Call) RunAndReturn(run func(context.Context, *proto.GetAttrRequest, ...grpc.CallOption) (*proto.GetAttrReply, error)) *MockRpcFsClient_GetAttr_Call {
	_c.Call.Return(run)
	return _c
}

// GetXAttr provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) GetXAttr(ctx context.Context, in *proto.GetXAttrRequest, opts ...grpc.CallOption) (*proto.GetXAttrReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetXAttr")
	}

	var r0 *proto.GetXAttrReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetXAttrRequest, ...grpc.CallOption) (*proto.GetXAttrReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetXAttrRequest, ...grpc.CallOption) *proto.GetXAttrReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetXAttrReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetXAttrRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_GetXAttr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetXAttr'
type MockRpcFsClient_GetXAttr_Call struct {
	*mock.Call
}

// GetXAttr is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.GetXAttrRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) GetXAttr(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_GetXAttr_Call {
	return &MockRpcFsClient_GetXAttr_Call{Call: _e.mock.On("GetXAttr",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_GetXAttr_Call) Run(run func(ctx context.Context, in *proto.GetXAttrRequest, opts ...grpc.CallOption)) *MockRpcFsClient_GetXAttr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.GetXAttrRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_GetXAttr_Call) Return(_a0 *proto.GetXAttrReply, _a1 error) *MockRpcFsClient_GetXAttr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_GetXAttr_Call) RunAndReturn(run func(context.Context, *proto.GetXAttrRequest, ...grpc.CallOption) (*proto.GetXAttrReply, error)) *MockRpcFsClient_GetXAttr_Call {
	_c.Call.Return(run)
	return _c
}

// Mkdir provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Mkdir(ctx context.Context, in *proto.MkdirRequest, opts ...grpc.CallOption) (*proto.MkdirReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Mkdir")
	}

	var r0 *proto.MkdirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MkdirRequest, ...grpc.CallOption) (*proto.MkdirReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MkdirRequest, ...grpc.CallOption) *proto.MkdirReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.MkdirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.MkdirRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Mkdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mkdir'
type MockRpcFsClient_Mkdir_Call struct {
	*mock.Call
}

// Mkdir is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.MkdirRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Mkdir(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Mkdir_Call {
	return &MockRpcFsClient_Mkdir_Call{Call: _e.mock.On("Mkdir",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Mkdir_Call) Run(run func(ctx context.Context, in *proto.MkdirRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Mkdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.MkdirRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Mkdir_Call) Return(_a0 *proto.MkdirReply, _a1 error) *MockRpcFsClient_Mkdir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Mkdir_Call) RunAndReturn(run func(context.Context, *proto.MkdirRequest, ...grpc.CallOption) (*proto.MkdirReply, error)) *MockRpcFsClient_Mkdir_Call {
	_c.Call.Return(run)
	return _c
}

// OpenDir provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) OpenDir(ctx context.Context, in *proto.OpenDirRequest, opts ...grpc.CallOption) (*proto.OpenDirReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for OpenDir")
	}

	var r0 *proto.OpenDirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenDirRequest, ...grpc.CallOption) (*proto.OpenDirReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenDirRequest, ...grpc.CallOption) *proto.OpenDirReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.OpenDirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.OpenDirRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_OpenDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OpenDir'
type MockRpcFsClient_OpenDir_Call struct {
	*mock.Call
}

// OpenDir is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.OpenDirRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) OpenDir(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_OpenDir_Call {
	return &MockRpcFsClient_OpenDir_Call{Call: _e.mock.On("OpenDir",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_OpenDir_Call) Run(run func(ctx context.Context, in *proto.OpenDirRequest, opts ...grpc.CallOption)) *MockRpcFsClient_OpenDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.OpenDirRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_OpenDir_Call) Return(_a0 *proto.OpenDirReply, _a1 error) *MockRpcFsClient_OpenDir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_OpenDir_Call) RunAndReturn(run func(context.Context, *proto.OpenDirRequest, ...grpc.CallOption) (*proto.OpenDirReply, error)) *MockRpcFsClient_OpenDir_Call {
	_c.Call.Return(run)
	return _c
}

// Rename provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Rename(ctx context.Context, in *proto.RenameRequest, opts ...grpc.CallOption) (*proto.RenameReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 *proto.RenameReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RenameRequest, ...grpc.CallOption) (*proto.RenameReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RenameRequest, ...grpc.CallOption) *proto.RenameReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RenameReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.RenameRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Rename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rename'
type MockRpcFsClient_Rename_Call struct {
	*mock.Call
}

// Rename is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.RenameRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Rename(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Rename_Call {
	return &MockRpcFsClient_Rename_Call{Call: _e.mock.On("Rename",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Rename_Call) Run(run func(ctx context.Context, in *proto.RenameRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Rename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.RenameRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Rename_Call) Return(_a0 *proto.RenameReply, _a1 error) *MockRpcFsClient_Rename_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Rename_Call) RunAndReturn(run func(context.Context, *proto.RenameRequest, ...grpc.CallOption) (*proto.RenameReply, error)) *MockRpcFsClient_Rename_Call {
	_c.Call.Return(run)
	return _c
}

// Rmdir provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Rmdir(ctx context.Context, in *proto.RmdirRequest, opts ...grpc.CallOption) (*proto.RmdirReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Rmdir")
	}

	var r0 *proto.RmdirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RmdirRequest, ...grpc.CallOption) (*proto.RmdirReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RmdirRequest, ...grpc.CallOption) *proto.RmdirReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RmdirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.RmdirRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Rmdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rmdir'
type MockRpcFsClient_Rmdir_Call struct {
	*mock.Call
}

// Rmdir is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.RmdirRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Rmdir(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Rmdir_Call {
	return &MockRpcFsClient_Rmdir_Call{Call: _e.mock.On("Rmdir",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Rmdir_Call) Run(run func(ctx context.Context, in *proto.RmdirRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Rmdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.RmdirRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Rmdir_Call) Return(_a0 *proto.RmdirReply, _a1 error) *MockRpcFsClient_Rmdir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Rmdir_Call) RunAndReturn(run func(context.Context, *proto.RmdirRequest, ...grpc.CallOption) (*proto.RmdirReply, error)) *MockRpcFsClient_Rmdir_Call {
	_c.Call.Return(run)
	return _c
}

// StatFs provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) StatFs(ctx context.Context, in *proto.StatFsRequest, opts ...grpc.CallOption) (*proto.StatFsReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StatFs")
	}

	var r0 *proto.StatFsReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.StatFsRequest, ...grpc.CallOption) (*proto.StatFsReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.StatFsRequest, ...grpc.CallOption) *proto.StatFsReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.StatFsReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.StatFsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_StatFs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatFs'
type MockRpcFsClient_StatFs_Call struct {
	*mock.Call
}

// StatFs is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.StatFsRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) StatFs(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_StatFs_Call {
	return &MockRpcFsClient_StatFs_Call{Call: _e.mock.On("StatFs",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_StatFs_Call) Run(run func(ctx context.Context, in *proto.StatFsRequest, opts ...grpc.CallOption)) *MockRpcFsClient_StatFs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.StatFsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_StatFs_Call) Return(_a0 *proto.StatFsReply, _a1 error) *MockRpcFsClient_StatFs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_StatFs_Call) RunAndReturn(run func(context.Context, *proto.StatFsRequest, ...grpc.CallOption) (*proto.StatFsReply, error)) *MockRpcFsClient_StatFs_Call {
	_c.Call.Return(run)
	return _c
}

// Truncate provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Truncate(ctx context.Context, in *proto.TruncateRequest, opts ...grpc.CallOption) (*proto.TruncateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Truncate")
	}

	var r0 *proto.TruncateReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.TruncateRequest, ...grpc.CallOption) (*proto.TruncateReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.TruncateRequest, ...grpc.CallOption) *proto.TruncateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.TruncateReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.TruncateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Truncate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Truncate'
type MockRpcFsClient_Truncate_Call struct {
	*mock.Call
}

// Truncate is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.TruncateRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Truncate(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Truncate_Call {
	return &MockRpcFsClient_Truncate_Call{Call: _e.mock.On("Truncate",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Truncate_Call) Run(run func(ctx context.Context, in *proto.TruncateRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Truncate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.TruncateRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Truncate_Call) Return(_a0 *proto.TruncateReply, _a1 error) *MockRpcFsClient_Truncate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Truncate_Call) RunAndReturn(run func(context.Context, *proto.TruncateRequest, ...grpc.CallOption) (*proto.TruncateReply, error)) *MockRpcFsClient_Truncate_Call {
	_c.Call.Return(run)
	return _c
}

// Unlink provides a mock function with given fields: ctx, in, opts
func (_m *MockRpcFsClient) Unlink(ctx context.Context, in *proto.UnlinkRequest, opts ...grpc.CallOption) (*proto.UnlinkReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Unlink")
	}

	var r0 *proto.UnlinkReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UnlinkRequest, ...grpc.CallOption) (*proto.UnlinkReply, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UnlinkRequest, ...grpc.CallOption) *proto.UnlinkReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UnlinkReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.UnlinkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsClient_Unlink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlink'
type MockRpcFsClient_Unlink_Call struct {
	*mock.Call
}

// Unlink is a helper method to define mock.On call
//   - ctx context.Context
//   - in *proto.UnlinkRequest
//   - opts ...grpc.CallOption
func (_e *MockRpcFsClient_Expecter) Unlink(ctx interface{}, in interface{}, opts ...interface{}) *MockRpcFsClient_Unlink_Call {
	return &MockRpcFsClient_Unlink_Call{Call: _e.mock.On("Unlink",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockRpcFsClient_Unlink_Call) Run(run func(ctx context.Context, in *proto.UnlinkRequest, opts ...grpc.CallOption)) *MockRpcFsClient_Unlink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*proto.UnlinkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockRpcFsClient_Unlink_Call) Return(_a0 *proto.UnlinkReply, _a1 error) *MockRpcFsClient_Unlink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsClient_Unlink_Call) RunAndReturn(run func(context.Context, *proto.UnlinkRequest, ...grpc.CallOption) (*proto.UnlinkReply, error)) *MockRpcFsClient_Unlink_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRpcFsClient creates a new instance of MockRpcFsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRpcFsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRpcFsClient {
	mock := &MockRpcFsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
