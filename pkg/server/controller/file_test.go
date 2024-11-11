package controller

import (
	"context"
	"gmountie/internal/mocks/pkg/server/service"
	"testing"

	"gmountie/pkg/proto"

	"github.com/hanwen/go-fuse/v2/fuse"
	"github.com/hanwen/go-fuse/v2/fuse/nodefs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	nodefs2 "gmountie/internal/mocks/github.com/hanwen/go-fuse/v2/fuse/nodefs"
	pathfs2 "gmountie/internal/mocks/github.com/hanwen/go-fuse/v2/fuse/pathfs"
)

type RpcFileServerTestSuite struct {
	suite.Suite
	server    *RpcFileServerImpl
	fsService *service.MockVolumeService
}

func (s *RpcFileServerTestSuite) SetupTest() {
	s.fsService = new(service.MockVolumeService)
	s.server = NewRpcFileServer(s.fsService)
}

func (s *RpcFileServerTestSuite) TestOpen() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	ctx := context.Background()
	mockFs.EXPECT().Open("/test/path", uint32(0), mock.Anything).Return(nodefs.NewDefaultFile(), fuse.OK)

	// Test.
	request := &proto.OpenRequest{Volume: "testVolume", Path: "/test/path", Flags: 0, Caller: CreateCaller(0, 0, 0)}
	reply, err := s.server.Open(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

func (s *RpcFileServerTestSuite) TestCreate() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	ctx := context.Background()
	mockFs.EXPECT().Create("/test/path", uint32(0), uint32(0), mock.Anything).Return(nodefs.NewDefaultFile(), fuse.OK)

	// Test.
	request := &proto.CreateRequest{Volume: "testVolume", Path: "/test/path", Flags: 0, Mode: 0, Caller: CreateCaller(0, 0, 0)}
	reply, err := s.server.Create(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

func (s *RpcFileServerTestSuite) TestRead() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	mockFile := new(nodefs2.MockFile)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	s.server.addFile(1, "/test/path", mockFile)
	ctx := context.Background()
	mockFile.EXPECT().Read(mock.Anything, int64(0)).Return(fuse.ReadResultData([]byte("test data")), fuse.OK)

	// Test.
	request := &proto.ReadRequest{Fd: 1, Size: 1024, Offset: 0}
	reply, err := s.server.Read(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

func (s *RpcFileServerTestSuite) TestWrite() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	mockFile := new(nodefs2.MockFile)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	s.server.addFile(1, "/test/path", mockFile)
	ctx := context.Background()
	mockFile.EXPECT().Write([]byte("test data"), int64(0)).Return(uint32(len("test data")), fuse.OK)

	// Test.
	request := &proto.WriteRequest{Fd: 1, Bytes: []byte("test data"), Offset: 0}
	reply, err := s.server.Write(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

func (s *RpcFileServerTestSuite) TestFsync() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	mockFile := new(nodefs2.MockFile)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	s.server.addFile(1, "/test/path", mockFile)
	ctx := context.Background()
	mockFile.EXPECT().Fsync(int(0)).Return(fuse.OK)

	// Test.
	request := &proto.FsyncRequest{Fd: 1, Flags: 0}
	reply, err := s.server.Fsync(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

func (s *RpcFileServerTestSuite) TestRelease() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	mockFile := new(nodefs2.MockFile)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	s.server.addFile(1, "/test/path", mockFile)
	ctx := context.Background()
	mockFile.EXPECT().Release().Return()

	// Test.
	request := &proto.ReleaseRequest{Fd: 1}
	reply, err := s.server.Release(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
}

func (s *RpcFileServerTestSuite) TestFlush() {
	// Setup.
	mockFs := new(pathfs2.MockFileSystem)
	mockFile := new(nodefs2.MockFile)
	s.fsService.On("GetVolumeFileSystem", "testVolume").Return(mockFs, nil)
	s.server.addFile(1, "/test/path", mockFile)
	ctx := context.Background()
	mockFile.EXPECT().Flush().Return(fuse.OK)

	// Test.
	request := &proto.FlushRequest{Fd: 1}
	reply, err := s.server.Flush(ctx, request)

	// Verify.
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Equal(int32(fuse.OK), reply.Status)
}

// --------- Helper Functions ---------

// CreateCaller creates a caller object
func CreateCaller(uid, gid, pid uint32) *proto.Caller {
	return &proto.Caller{
		Owner: &proto.Owner{
			Uid: uid,
			Gid: gid,
		},
		Pid: pid,
	}
}

func TestRpcFileServerTestSuite(t *testing.T) {
	suite.Run(t, new(RpcFileServerTestSuite))
}
