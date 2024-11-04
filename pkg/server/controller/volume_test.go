package controller

import (
	"context"
	"gmountie/mocks/pkg/server/service"
	"gmountie/pkg/common"
	"gmountie/pkg/proto"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type VolumeServiceTestSuite struct {
	suite.Suite
	server  *VolumeServiceImpl
	service *service.MockVolumeService
}

func (s *VolumeServiceTestSuite) SetupTest() {
	s.service = new(service.MockVolumeService)
	s.server = NewVolumeService(s.service)
}

func (s *VolumeServiceTestSuite) TestList_Success() {
	// Setup
	expectedVolumes := []common.Volume{
		{Name: "volume1"},
		{Name: "volume2"},
	}
	s.service.On("List").Return(expectedVolumes, nil)

	// Test
	reply, err := s.server.List(context.Background(), &proto.VolumeListRequest{})

	// Verify
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Len(reply.Volumes, 2)
	s.Assert().Equal("volume1", reply.Volumes[0].Name)
	s.Assert().Equal("volume2", reply.Volumes[1].Name)
	s.service.AssertExpectations(s.T())
}

func (s *VolumeServiceTestSuite) TestList_EmptyList() {
	// Setup
	s.service.On("List").Return([]common.Volume{}, nil)

	// Test
	reply, err := s.server.List(context.Background(), &proto.VolumeListRequest{})

	// Verify
	s.Require().NoError(err)
	s.Assert().NotNil(reply)
	s.Assert().Empty(reply.Volumes)
	s.service.AssertExpectations(s.T())
}

func (s *VolumeServiceTestSuite) TestList_ServiceError() {
	// Setup
	expectedError := errors.New("test")
	s.service.On("List").Return(nil, expectedError)

	// Test
	reply, err := s.server.List(context.Background(), &proto.VolumeListRequest{})

	// Verify
	s.Require().Error(err)
	s.Assert().Nil(reply)
	s.service.AssertExpectations(s.T())
}

func TestVolumeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeServiceTestSuite))
}
