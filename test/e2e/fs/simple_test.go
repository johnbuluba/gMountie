package fs

import (
	"gmountie/test/e2e/utils"
	"os"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/suite"
)

type SimpleFSTestSuite struct {
	suite.Suite
	testAppCtx *utils.AppTestingContext
}

func (s *SimpleFSTestSuite) SetupSuite() {
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

func (s *SimpleFSTestSuite) TestFS() {
	// Setup.
	// Get the volume.
	volume := s.testAppCtx.GetVolumes()[0]
	s.Require().NotNil(volume)
	s.Require().GreaterOrEqual(len(volume.GeneratedFiles), 1)
	// Mount the volume.
	err := s.testAppCtx.MountVolume(volume)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		// Unmount the volume.
		err := s.testAppCtx.GetClientApp().MounterService.Unmount(volume.Name)
		s.Require().NoError(err)
	})

	// Test.
	testFS := os.DirFS(volume.GetMountPath())
	err = fstest.TestFS(testFS, volume.GeneratedFiles...)

	// Assert.
	s.Require().NoError(err)
}

func (s *SimpleFSTestSuite) TearDownSuite() {
	err := s.testAppCtx.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

func TestSimpleFSTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleFSTestSuite))
}
