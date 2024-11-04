package fs

import (
	"gmountie/test/e2e/utils"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/suite"
)

type SimpleFSTestSuite struct {
	suite.Suite
	testAppCtx *utils.AppTestingContext
	volume     *utils.TestVolume
}

func (s *SimpleFSTestSuite) SetupSuite() {
	testAppCtx, err := utils.NewAppTestingContext(
		utils.WithBasicAuth("test", "test"),
		utils.WithRandomTestVolume(true),
	)
	if err != nil {
		s.T().Fatal(err)
	}
	utils.Must0(s.T(), testAppCtx.Start())
	s.testAppCtx = testAppCtx
	// Mount the volume.
	s.volume = s.testAppCtx.GetVolumes()[0]
	s.Require().NotNil(s.volume)
	s.Require().GreaterOrEqual(len(s.volume.GeneratedFiles), 1)
	s.testAppCtx.MountVolume(s.volume)
}

func (s *SimpleFSTestSuite) TearDownSuite() {
	err := s.testAppCtx.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *SimpleFSTestSuite) TestFS() {
	// Test.
	testFS := os.DirFS(s.volume.GetMountPath())
	err := fstest.TestFS(testFS, s.volume.GeneratedFiles...)

	// Assert.
	s.Require().NoError(err)
}

func (s *SimpleFSTestSuite) Test_CopyDelete() {
	// Setup.
	// Create a destination directory.
	destDir := filepath.Join(s.volume.GetMountPath(), "test")
	s.Require().NoError(os.MkdirAll(destDir, 0755))
	testFS := os.DirFS(s.volume.GetRandomFilesPath())

	// Test.
	// Copy the test files to the destination directory.
	s.Require().NoError(os.CopyFS(destDir, testFS))
	// Delete the destination directory.
	s.Require().NoError(os.RemoveAll(destDir))
}

func TestSimpleFSTestSuite(t *testing.T) {
	suite.Run(t, new(SimpleFSTestSuite))
}
