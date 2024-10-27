package api

import (
	"context"
	"gmountie/test/e2e/utils"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RpcFileServerTestSuite struct {
	suite.Suite
	testAppCtx *utils.AppTestingContext
}

func (s *RpcFileServerTestSuite) SetupSuite() {
	// Create a new auth service.
	testAppCtx, err := utils.NewAppTestingContext(
		utils.WithBasicAuth("test", "test"),
		utils.WithRandomTestVolume(true),
	)
	if err != nil {
		s.T().Fatal(err)
	}
	err = testAppCtx.Start()
	if err != nil {
		s.T().Fatal(err)
	}
	s.testAppCtx = testAppCtx
}

func (s *RpcFileServerTestSuite) TestListFiles() {
	clientVolumes, err := s.testAppCtx.GetClientApp().VolumeService.GetVolumes(context.Background())
	serverVolumes, err := s.testAppCtx.GetServerApp().VolumeService.List()

	s.Assert().NoError(err)
	s.Assert().Equal(clientVolumes, serverVolumes)
}

func (s *RpcFileServerTestSuite) TearDownSuite() {
	err := s.testAppCtx.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestRpcFileServerTestSuite(t *testing.T) {
	suite.Run(t, new(RpcFileServerTestSuite))
}
