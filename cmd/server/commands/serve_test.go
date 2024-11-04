package commands

import (
	"bytes"
	"gmountie/pkg/server/config"
	"gmountie/test/e2e/utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type ServeCmdTestSuite struct {
	suite.Suite
	cmd     *cobra.Command
	buf     *bytes.Buffer
	tempDir string
}

func (s *ServeCmdTestSuite) SetupTest() {
	s.tempDir, _ = os.MkdirTemp("", "serveCmd_test")
	utils.Must0(s.T(), os.Setenv("HOME", s.tempDir))

	s.cmd = &cobra.Command{Use: "root"}
	s.cmd.AddCommand(serveCmd)
	s.buf = new(bytes.Buffer)
	s.cmd.SetOutput(s.buf)
}

func (s *ServeCmdTestSuite) TearDownTest() {
	utils.Must0(s.T(), os.RemoveAll(s.tempDir))
}

func (s *ServeCmdTestSuite) TestServeCmd_ExecuteWithoutConfig() {
	// Setup
	originalServerStart := serverStart
	defer func() { serverStart = originalServerStart }()

	serverStartCalled := false
	serverStart = func(cfg *config.Config) error {
		serverStartCalled = true
		return nil
	}

	// Test
	s.cmd.SetArgs([]string{"serve"})
	err := s.cmd.Execute()

	// Verify
	s.Require().NoError(err)
	s.Assert().True(serverStartCalled)

	// Check if default config was created
	defaultConfigPath := config.GetDefaultConfigPath()
	_, err = os.Stat(defaultConfigPath)
	s.Assert().NoError(err)
}

func (s *ServeCmdTestSuite) TestServeCmd_ExecuteWithInvalidConfig() {
	// Setup
	configFile := filepath.Join(s.tempDir, ".config", "gmountie", "config.yaml")
	utils.Must0(s.T(), os.MkdirAll(filepath.Dir(configFile), 0755))
	utils.Must0(s.T(), os.WriteFile(configFile, []byte("test-config"), 0644))

	originalServerStart := serverStart
	defer func() { serverStart = originalServerStart }()

	// Test
	s.cmd.SetArgs([]string{"serve"})
	err := s.cmd.Execute()

	// Verify
	s.Require().Error(err)
}

func TestServeCmdSuite(t *testing.T) {
	suite.Run(t, new(ServeCmdTestSuite))
}
