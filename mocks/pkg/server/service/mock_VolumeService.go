// Code generated by mockery v2.46.3. DO NOT EDIT.

package service

import (
	common "gmountie/pkg/common"

	mock "github.com/stretchr/testify/mock"

	pathfs "github.com/hanwen/go-fuse/v2/fuse/pathfs"
)

// MockVolumeService is an autogenerated mock type for the VolumeService type
type MockVolumeService struct {
	mock.Mock
}

type MockVolumeService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockVolumeService) EXPECT() *MockVolumeService_Expecter {
	return &MockVolumeService_Expecter{mock: &_m.Mock}
}

// GetVolumeFileSystem provides a mock function with given fields: name
func (_m *MockVolumeService) GetVolumeFileSystem(name string) (pathfs.FileSystem, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetVolumeFileSystem")
	}

	var r0 pathfs.FileSystem
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (pathfs.FileSystem, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) pathfs.FileSystem); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pathfs.FileSystem)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumeService_GetVolumeFileSystem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVolumeFileSystem'
type MockVolumeService_GetVolumeFileSystem_Call struct {
	*mock.Call
}

// GetVolumeFileSystem is a helper method to define mock.On call
//   - name string
func (_e *MockVolumeService_Expecter) GetVolumeFileSystem(name interface{}) *MockVolumeService_GetVolumeFileSystem_Call {
	return &MockVolumeService_GetVolumeFileSystem_Call{Call: _e.mock.On("GetVolumeFileSystem", name)}
}

func (_c *MockVolumeService_GetVolumeFileSystem_Call) Run(run func(name string)) *MockVolumeService_GetVolumeFileSystem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockVolumeService_GetVolumeFileSystem_Call) Return(_a0 pathfs.FileSystem, _a1 error) *MockVolumeService_GetVolumeFileSystem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumeService_GetVolumeFileSystem_Call) RunAndReturn(run func(string) (pathfs.FileSystem, error)) *MockVolumeService_GetVolumeFileSystem_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields:
func (_m *MockVolumeService) List() ([]common.Volume, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []common.Volume
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]common.Volume, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []common.Volume); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Volume)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockVolumeService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockVolumeService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
func (_e *MockVolumeService_Expecter) List() *MockVolumeService_List_Call {
	return &MockVolumeService_List_Call{Call: _e.mock.On("List")}
}

func (_c *MockVolumeService_List_Call) Run(run func()) *MockVolumeService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockVolumeService_List_Call) Return(_a0 []common.Volume, _a1 error) *MockVolumeService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockVolumeService_List_Call) RunAndReturn(run func() ([]common.Volume, error)) *MockVolumeService_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockVolumeService creates a new instance of MockVolumeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockVolumeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockVolumeService {
	mock := &MockVolumeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
