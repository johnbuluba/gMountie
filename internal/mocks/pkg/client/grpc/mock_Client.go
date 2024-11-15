// Code generated by mockery v2.46.3. DO NOT EDIT.

package grpc

import (
	proto "gmountie/pkg/proto"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockClient) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockClient_Expecter) Close() *MockClient_Close_Call {
	return &MockClient_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockClient_Close_Call) Run(run func()) *MockClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_Close_Call) Return(_a0 error) *MockClient_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Close_Call) RunAndReturn(run func() error) *MockClient_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockClient) Connect() {
	_m.Called()
}

// MockClient_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockClient_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockClient_Expecter) Connect() *MockClient_Connect_Call {
	return &MockClient_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockClient_Connect_Call) Run(run func()) *MockClient_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_Connect_Call) Return() *MockClient_Connect_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockClient_Connect_Call) RunAndReturn(run func()) *MockClient_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// File provides a mock function with given fields:
func (_m *MockClient) File() proto.RpcFileClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for File")
	}

	var r0 proto.RpcFileClient
	if rf, ok := ret.Get(0).(func() proto.RpcFileClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proto.RpcFileClient)
		}
	}

	return r0
}

// MockClient_File_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'File'
type MockClient_File_Call struct {
	*mock.Call
}

// File is a helper method to define mock.On call
func (_e *MockClient_Expecter) File() *MockClient_File_Call {
	return &MockClient_File_Call{Call: _e.mock.On("File")}
}

func (_c *MockClient_File_Call) Run(run func()) *MockClient_File_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_File_Call) Return(_a0 proto.RpcFileClient) *MockClient_File_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_File_Call) RunAndReturn(run func() proto.RpcFileClient) *MockClient_File_Call {
	_c.Call.Return(run)
	return _c
}

// Fs provides a mock function with given fields:
func (_m *MockClient) Fs() proto.RpcFsClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fs")
	}

	var r0 proto.RpcFsClient
	if rf, ok := ret.Get(0).(func() proto.RpcFsClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proto.RpcFsClient)
		}
	}

	return r0
}

// MockClient_Fs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fs'
type MockClient_Fs_Call struct {
	*mock.Call
}

// Fs is a helper method to define mock.On call
func (_e *MockClient_Expecter) Fs() *MockClient_Fs_Call {
	return &MockClient_Fs_Call{Call: _e.mock.On("Fs")}
}

func (_c *MockClient_Fs_Call) Run(run func()) *MockClient_Fs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_Fs_Call) Return(_a0 proto.RpcFsClient) *MockClient_Fs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Fs_Call) RunAndReturn(run func() proto.RpcFsClient) *MockClient_Fs_Call {
	_c.Call.Return(run)
	return _c
}

// GetEndpoint provides a mock function with given fields:
func (_m *MockClient) GetEndpoint() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEndpoint")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClient_GetEndpoint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEndpoint'
type MockClient_GetEndpoint_Call struct {
	*mock.Call
}

// GetEndpoint is a helper method to define mock.On call
func (_e *MockClient_Expecter) GetEndpoint() *MockClient_GetEndpoint_Call {
	return &MockClient_GetEndpoint_Call{Call: _e.mock.On("GetEndpoint")}
}

func (_c *MockClient_GetEndpoint_Call) Run(run func()) *MockClient_GetEndpoint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_GetEndpoint_Call) Return(_a0 string) *MockClient_GetEndpoint_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_GetEndpoint_Call) RunAndReturn(run func() string) *MockClient_GetEndpoint_Call {
	_c.Call.Return(run)
	return _c
}

// Volume provides a mock function with given fields:
func (_m *MockClient) Volume() proto.VolumeServiceClient {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Volume")
	}

	var r0 proto.VolumeServiceClient
	if rf, ok := ret.Get(0).(func() proto.VolumeServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proto.VolumeServiceClient)
		}
	}

	return r0
}

// MockClient_Volume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Volume'
type MockClient_Volume_Call struct {
	*mock.Call
}

// Volume is a helper method to define mock.On call
func (_e *MockClient_Expecter) Volume() *MockClient_Volume_Call {
	return &MockClient_Volume_Call{Call: _e.mock.On("Volume")}
}

func (_c *MockClient_Volume_Call) Run(run func()) *MockClient_Volume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClient_Volume_Call) Return(_a0 proto.VolumeServiceClient) *MockClient_Volume_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Volume_Call) RunAndReturn(run func() proto.VolumeServiceClient) *MockClient_Volume_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
