package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PathsTestSuite struct {
	suite.Suite
	tempDir string
}

func (s *PathsTestSuite) SetupTest() {
	var err error
	s.tempDir, err = os.MkdirTemp("", "paths_test")
	s.Require().NoError(err)
}

func (s *PathsTestSuite) TearDownTest() {
	os.RemoveAll(s.tempDir)
}

func (s *PathsTestSuite) TestGetDefaultConfigDir_WithXDG() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)

	// Test
	result := GetDefaultConfigDir()

	// Verify
	expected := filepath.Join(s.tempDir, DefaultConfigDirName)
	s.Assert().Equal(expected, result)
}

func (s *PathsTestSuite) TestGetDefaultConfigDir_WithoutXDG() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", "")
	s.T().Setenv("HOME", s.tempDir)

	// Test
	result := GetDefaultConfigDir()

	// Verify
	expected := filepath.Join(s.tempDir, ".config", DefaultConfigDirName)
	s.Assert().Equal(expected, result)
}

func (s *PathsTestSuite) TestGetDefaultConfigPath() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)

	// Test
	result := GetDefaultConfigPath()

	// Verify
	expected := filepath.Join(s.tempDir, DefaultConfigDirName, DefaultConfigFileName)
	s.Assert().Equal(expected, result)
}

func (s *PathsTestSuite) TestEnsureConfigDir() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)
	configDir := GetDefaultConfigDir()

	// Test
	err := EnsureConfigDir()

	// Verify
	s.Require().NoError(err)
	info, err := os.Stat(configDir)
	s.Require().NoError(err)
	s.Assert().True(info.IsDir())
}

func (s *PathsTestSuite) TestEnsureConfigDir_AlreadyExists() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)
	configDir := GetDefaultConfigDir()
	err := os.MkdirAll(configDir, 0755)
	s.Require().NoError(err)

	// Test
	err = EnsureConfigDir()

	// Verify
	s.Require().NoError(err)
	info, err := os.Stat(configDir)
	s.Require().NoError(err)
	s.Assert().True(info.IsDir())
}

func (s *PathsTestSuite) TestWriteDefaultConfig() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)
	err := EnsureConfigDir()
	s.Require().NoError(err)

	// Test
	err = WriteDefaultConfig()

	// Verify
	s.Require().NoError(err)
	configPath := GetDefaultConfigPath()
	content, err := os.ReadFile(configPath)
	s.Require().NoError(err)
	s.Assert().Equal(DefaultConfig, string(content))
}

func (s *PathsTestSuite) TestWriteDefaultConfig_OverwriteExisting() {
	// Setup
	s.T().Setenv("XDG_CONFIG_HOME", s.tempDir)
	err := EnsureConfigDir()
	s.Require().NoError(err)
	configPath := GetDefaultConfigPath()
	err = os.WriteFile(configPath, []byte("old content"), 0644)
	s.Require().NoError(err)

	// Test
	err = WriteDefaultConfig()

	// Verify
	s.Require().NoError(err)
	content, err := os.ReadFile(configPath)
	s.Require().NoError(err)
	s.Assert().Equal(DefaultConfig, string(content))
}

func TestPathsTestSuite(t *testing.T) {
	suite.Run(t, new(PathsTestSuite))
}
