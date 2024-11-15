// Code generated by mockery v2.46.3. DO NOT EDIT.

package proto

import (
	context "context"
	proto "gmountie/pkg/proto"

	mock "github.com/stretchr/testify/mock"
)

// MockRpcFsServer is an autogenerated mock type for the RpcFsServer type
type MockRpcFsServer struct {
	mock.Mock
}

type MockRpcFsServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRpcFsServer) EXPECT() *MockRpcFsServer_Expecter {
	return &MockRpcFsServer_Expecter{mock: &_m.Mock}
}

// Access provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Access(_a0 context.Context, _a1 *proto.AccessRequest) (*proto.AccessReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Access")
	}

	var r0 *proto.AccessReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AccessRequest) (*proto.AccessReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.AccessRequest) *proto.AccessReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.AccessReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.AccessRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Access_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Access'
type MockRpcFsServer_Access_Call struct {
	*mock.Call
}

// Access is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.AccessRequest
func (_e *MockRpcFsServer_Expecter) Access(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Access_Call {
	return &MockRpcFsServer_Access_Call{Call: _e.mock.On("Access", _a0, _a1)}
}

func (_c *MockRpcFsServer_Access_Call) Run(run func(_a0 context.Context, _a1 *proto.AccessRequest)) *MockRpcFsServer_Access_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.AccessRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Access_Call) Return(_a0 *proto.AccessReply, _a1 error) *MockRpcFsServer_Access_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Access_Call) RunAndReturn(run func(context.Context, *proto.AccessRequest) (*proto.AccessReply, error)) *MockRpcFsServer_Access_Call {
	_c.Call.Return(run)
	return _c
}

// Chmod provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Chmod(_a0 context.Context, _a1 *proto.ChmodRequest) (*proto.ChmodReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Chmod")
	}

	var r0 *proto.ChmodReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChmodRequest) (*proto.ChmodReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChmodRequest) *proto.ChmodReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ChmodReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ChmodRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Chmod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chmod'
type MockRpcFsServer_Chmod_Call struct {
	*mock.Call
}

// Chmod is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.ChmodRequest
func (_e *MockRpcFsServer_Expecter) Chmod(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Chmod_Call {
	return &MockRpcFsServer_Chmod_Call{Call: _e.mock.On("Chmod", _a0, _a1)}
}

func (_c *MockRpcFsServer_Chmod_Call) Run(run func(_a0 context.Context, _a1 *proto.ChmodRequest)) *MockRpcFsServer_Chmod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.ChmodRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Chmod_Call) Return(_a0 *proto.ChmodReply, _a1 error) *MockRpcFsServer_Chmod_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Chmod_Call) RunAndReturn(run func(context.Context, *proto.ChmodRequest) (*proto.ChmodReply, error)) *MockRpcFsServer_Chmod_Call {
	_c.Call.Return(run)
	return _c
}

// Chown provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Chown(_a0 context.Context, _a1 *proto.ChownRequest) (*proto.ChownReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Chown")
	}

	var r0 *proto.ChownReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChownRequest) (*proto.ChownReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ChownRequest) *proto.ChownReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ChownReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.ChownRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Chown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Chown'
type MockRpcFsServer_Chown_Call struct {
	*mock.Call
}

// Chown is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.ChownRequest
func (_e *MockRpcFsServer_Expecter) Chown(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Chown_Call {
	return &MockRpcFsServer_Chown_Call{Call: _e.mock.On("Chown", _a0, _a1)}
}

func (_c *MockRpcFsServer_Chown_Call) Run(run func(_a0 context.Context, _a1 *proto.ChownRequest)) *MockRpcFsServer_Chown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.ChownRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Chown_Call) Return(_a0 *proto.ChownReply, _a1 error) *MockRpcFsServer_Chown_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Chown_Call) RunAndReturn(run func(context.Context, *proto.ChownRequest) (*proto.ChownReply, error)) *MockRpcFsServer_Chown_Call {
	_c.Call.Return(run)
	return _c
}

// GetAttr provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) GetAttr(_a0 context.Context, _a1 *proto.GetAttrRequest) (*proto.GetAttrReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAttr")
	}

	var r0 *proto.GetAttrReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetAttrRequest) (*proto.GetAttrReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetAttrRequest) *proto.GetAttrReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetAttrReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetAttrRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_GetAttr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAttr'
type MockRpcFsServer_GetAttr_Call struct {
	*mock.Call
}

// GetAttr is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.GetAttrRequest
func (_e *MockRpcFsServer_Expecter) GetAttr(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_GetAttr_Call {
	return &MockRpcFsServer_GetAttr_Call{Call: _e.mock.On("GetAttr", _a0, _a1)}
}

func (_c *MockRpcFsServer_GetAttr_Call) Run(run func(_a0 context.Context, _a1 *proto.GetAttrRequest)) *MockRpcFsServer_GetAttr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.GetAttrRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_GetAttr_Call) Return(_a0 *proto.GetAttrReply, _a1 error) *MockRpcFsServer_GetAttr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_GetAttr_Call) RunAndReturn(run func(context.Context, *proto.GetAttrRequest) (*proto.GetAttrReply, error)) *MockRpcFsServer_GetAttr_Call {
	_c.Call.Return(run)
	return _c
}

// GetXAttr provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) GetXAttr(_a0 context.Context, _a1 *proto.GetXAttrRequest) (*proto.GetXAttrReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetXAttr")
	}

	var r0 *proto.GetXAttrReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetXAttrRequest) (*proto.GetXAttrReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetXAttrRequest) *proto.GetXAttrReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetXAttrReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetXAttrRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_GetXAttr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetXAttr'
type MockRpcFsServer_GetXAttr_Call struct {
	*mock.Call
}

// GetXAttr is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.GetXAttrRequest
func (_e *MockRpcFsServer_Expecter) GetXAttr(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_GetXAttr_Call {
	return &MockRpcFsServer_GetXAttr_Call{Call: _e.mock.On("GetXAttr", _a0, _a1)}
}

func (_c *MockRpcFsServer_GetXAttr_Call) Run(run func(_a0 context.Context, _a1 *proto.GetXAttrRequest)) *MockRpcFsServer_GetXAttr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.GetXAttrRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_GetXAttr_Call) Return(_a0 *proto.GetXAttrReply, _a1 error) *MockRpcFsServer_GetXAttr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_GetXAttr_Call) RunAndReturn(run func(context.Context, *proto.GetXAttrRequest) (*proto.GetXAttrReply, error)) *MockRpcFsServer_GetXAttr_Call {
	_c.Call.Return(run)
	return _c
}

// Mkdir provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Mkdir(_a0 context.Context, _a1 *proto.MkdirRequest) (*proto.MkdirReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Mkdir")
	}

	var r0 *proto.MkdirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MkdirRequest) (*proto.MkdirReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.MkdirRequest) *proto.MkdirReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.MkdirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.MkdirRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Mkdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mkdir'
type MockRpcFsServer_Mkdir_Call struct {
	*mock.Call
}

// Mkdir is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.MkdirRequest
func (_e *MockRpcFsServer_Expecter) Mkdir(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Mkdir_Call {
	return &MockRpcFsServer_Mkdir_Call{Call: _e.mock.On("Mkdir", _a0, _a1)}
}

func (_c *MockRpcFsServer_Mkdir_Call) Run(run func(_a0 context.Context, _a1 *proto.MkdirRequest)) *MockRpcFsServer_Mkdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.MkdirRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Mkdir_Call) Return(_a0 *proto.MkdirReply, _a1 error) *MockRpcFsServer_Mkdir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Mkdir_Call) RunAndReturn(run func(context.Context, *proto.MkdirRequest) (*proto.MkdirReply, error)) *MockRpcFsServer_Mkdir_Call {
	_c.Call.Return(run)
	return _c
}

// OpenDir provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) OpenDir(_a0 context.Context, _a1 *proto.OpenDirRequest) (*proto.OpenDirReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for OpenDir")
	}

	var r0 *proto.OpenDirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenDirRequest) (*proto.OpenDirReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.OpenDirRequest) *proto.OpenDirReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.OpenDirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.OpenDirRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_OpenDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OpenDir'
type MockRpcFsServer_OpenDir_Call struct {
	*mock.Call
}

// OpenDir is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.OpenDirRequest
func (_e *MockRpcFsServer_Expecter) OpenDir(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_OpenDir_Call {
	return &MockRpcFsServer_OpenDir_Call{Call: _e.mock.On("OpenDir", _a0, _a1)}
}

func (_c *MockRpcFsServer_OpenDir_Call) Run(run func(_a0 context.Context, _a1 *proto.OpenDirRequest)) *MockRpcFsServer_OpenDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.OpenDirRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_OpenDir_Call) Return(_a0 *proto.OpenDirReply, _a1 error) *MockRpcFsServer_OpenDir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_OpenDir_Call) RunAndReturn(run func(context.Context, *proto.OpenDirRequest) (*proto.OpenDirReply, error)) *MockRpcFsServer_OpenDir_Call {
	_c.Call.Return(run)
	return _c
}

// Rename provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Rename(_a0 context.Context, _a1 *proto.RenameRequest) (*proto.RenameReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Rename")
	}

	var r0 *proto.RenameReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RenameRequest) (*proto.RenameReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RenameRequest) *proto.RenameReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RenameReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.RenameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Rename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rename'
type MockRpcFsServer_Rename_Call struct {
	*mock.Call
}

// Rename is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.RenameRequest
func (_e *MockRpcFsServer_Expecter) Rename(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Rename_Call {
	return &MockRpcFsServer_Rename_Call{Call: _e.mock.On("Rename", _a0, _a1)}
}

func (_c *MockRpcFsServer_Rename_Call) Run(run func(_a0 context.Context, _a1 *proto.RenameRequest)) *MockRpcFsServer_Rename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.RenameRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Rename_Call) Return(_a0 *proto.RenameReply, _a1 error) *MockRpcFsServer_Rename_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Rename_Call) RunAndReturn(run func(context.Context, *proto.RenameRequest) (*proto.RenameReply, error)) *MockRpcFsServer_Rename_Call {
	_c.Call.Return(run)
	return _c
}

// Rmdir provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Rmdir(_a0 context.Context, _a1 *proto.RmdirRequest) (*proto.RmdirReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Rmdir")
	}

	var r0 *proto.RmdirReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RmdirRequest) (*proto.RmdirReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.RmdirRequest) *proto.RmdirReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.RmdirReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.RmdirRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Rmdir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rmdir'
type MockRpcFsServer_Rmdir_Call struct {
	*mock.Call
}

// Rmdir is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.RmdirRequest
func (_e *MockRpcFsServer_Expecter) Rmdir(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Rmdir_Call {
	return &MockRpcFsServer_Rmdir_Call{Call: _e.mock.On("Rmdir", _a0, _a1)}
}

func (_c *MockRpcFsServer_Rmdir_Call) Run(run func(_a0 context.Context, _a1 *proto.RmdirRequest)) *MockRpcFsServer_Rmdir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.RmdirRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Rmdir_Call) Return(_a0 *proto.RmdirReply, _a1 error) *MockRpcFsServer_Rmdir_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Rmdir_Call) RunAndReturn(run func(context.Context, *proto.RmdirRequest) (*proto.RmdirReply, error)) *MockRpcFsServer_Rmdir_Call {
	_c.Call.Return(run)
	return _c
}

// StatFs provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) StatFs(_a0 context.Context, _a1 *proto.StatFsRequest) (*proto.StatFsReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for StatFs")
	}

	var r0 *proto.StatFsReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.StatFsRequest) (*proto.StatFsReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.StatFsRequest) *proto.StatFsReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.StatFsReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.StatFsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_StatFs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatFs'
type MockRpcFsServer_StatFs_Call struct {
	*mock.Call
}

// StatFs is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.StatFsRequest
func (_e *MockRpcFsServer_Expecter) StatFs(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_StatFs_Call {
	return &MockRpcFsServer_StatFs_Call{Call: _e.mock.On("StatFs", _a0, _a1)}
}

func (_c *MockRpcFsServer_StatFs_Call) Run(run func(_a0 context.Context, _a1 *proto.StatFsRequest)) *MockRpcFsServer_StatFs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.StatFsRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_StatFs_Call) Return(_a0 *proto.StatFsReply, _a1 error) *MockRpcFsServer_StatFs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_StatFs_Call) RunAndReturn(run func(context.Context, *proto.StatFsRequest) (*proto.StatFsReply, error)) *MockRpcFsServer_StatFs_Call {
	_c.Call.Return(run)
	return _c
}

// Truncate provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Truncate(_a0 context.Context, _a1 *proto.TruncateRequest) (*proto.TruncateReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Truncate")
	}

	var r0 *proto.TruncateReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.TruncateRequest) (*proto.TruncateReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.TruncateRequest) *proto.TruncateReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.TruncateReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.TruncateRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Truncate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Truncate'
type MockRpcFsServer_Truncate_Call struct {
	*mock.Call
}

// Truncate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.TruncateRequest
func (_e *MockRpcFsServer_Expecter) Truncate(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Truncate_Call {
	return &MockRpcFsServer_Truncate_Call{Call: _e.mock.On("Truncate", _a0, _a1)}
}

func (_c *MockRpcFsServer_Truncate_Call) Run(run func(_a0 context.Context, _a1 *proto.TruncateRequest)) *MockRpcFsServer_Truncate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.TruncateRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Truncate_Call) Return(_a0 *proto.TruncateReply, _a1 error) *MockRpcFsServer_Truncate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Truncate_Call) RunAndReturn(run func(context.Context, *proto.TruncateRequest) (*proto.TruncateReply, error)) *MockRpcFsServer_Truncate_Call {
	_c.Call.Return(run)
	return _c
}

// Unlink provides a mock function with given fields: _a0, _a1
func (_m *MockRpcFsServer) Unlink(_a0 context.Context, _a1 *proto.UnlinkRequest) (*proto.UnlinkReply, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Unlink")
	}

	var r0 *proto.UnlinkReply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UnlinkRequest) (*proto.UnlinkReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UnlinkRequest) *proto.UnlinkReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UnlinkReply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.UnlinkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRpcFsServer_Unlink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlink'
type MockRpcFsServer_Unlink_Call struct {
	*mock.Call
}

// Unlink is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.UnlinkRequest
func (_e *MockRpcFsServer_Expecter) Unlink(_a0 interface{}, _a1 interface{}) *MockRpcFsServer_Unlink_Call {
	return &MockRpcFsServer_Unlink_Call{Call: _e.mock.On("Unlink", _a0, _a1)}
}

func (_c *MockRpcFsServer_Unlink_Call) Run(run func(_a0 context.Context, _a1 *proto.UnlinkRequest)) *MockRpcFsServer_Unlink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.UnlinkRequest))
	})
	return _c
}

func (_c *MockRpcFsServer_Unlink_Call) Return(_a0 *proto.UnlinkReply, _a1 error) *MockRpcFsServer_Unlink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRpcFsServer_Unlink_Call) RunAndReturn(run func(context.Context, *proto.UnlinkRequest) (*proto.UnlinkReply, error)) *MockRpcFsServer_Unlink_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedRpcFsServer provides a mock function with given fields:
func (_m *MockRpcFsServer) mustEmbedUnimplementedRpcFsServer() {
	_m.Called()
}

// MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedRpcFsServer'
type MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedRpcFsServer is a helper method to define mock.On call
func (_e *MockRpcFsServer_Expecter) mustEmbedUnimplementedRpcFsServer() *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call {
	return &MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call{Call: _e.mock.On("mustEmbedUnimplementedRpcFsServer")}
}

func (_c *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call) Run(run func()) *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call) Return() *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call) RunAndReturn(run func()) *MockRpcFsServer_mustEmbedUnimplementedRpcFsServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRpcFsServer creates a new instance of MockRpcFsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRpcFsServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRpcFsServer {
	mock := &MockRpcFsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
