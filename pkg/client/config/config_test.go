package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	fullConf    string
	minimalConf string
}

func (s *ConfigTestSuite) SetupTest() {
	s.fullConf = `
server:
  address: 0.0.0.0
  port: 9449
  tls: false
auth:
  type: none
`
	s.minimalConf = `
server:
  address: 127.0.0.1
auth:
  type: none
`
}

// Test String() function
func (s *ConfigTestSuite) TestString() {
	cfg, err := LoadConfigFromString(s.fullConf)
	s.Require().NoError(err)

	str, err := cfg.String()
	s.Require().NoError(err)
	s.Assert().Contains(str, "address: 0.0.0.0")
	s.Assert().Contains(str, "port: 9449")
	s.Assert().Contains(str, "type: none")
}

// Test Save() function
func (s *ConfigTestSuite) TestSave() {
	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "config_test_*.yaml")
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		os.Remove(tmpfile.Name())
	})

	// Load config and set path
	cfg, err := LoadConfigFromString(s.fullConf)
	s.Require().NoError(err)

	// Test Save
	err = cfg.Save(tmpfile.Name())
	s.Require().NoError(err)

	// Verify file contents
	content, err := os.ReadFile(tmpfile.Name())
	s.Require().NoError(err)
	s.Assert().Contains(string(content), "address: 0.0.0.0")
}

// Test Save() with invalid path
func (s *ConfigTestSuite) TestSaveInvalidPath() {
	path := "/nonexistent/path/config.yaml"
	cfg, err := LoadConfigFromString(s.fullConf)
	s.Require().NoError(err)

	err = cfg.Save(path)
	s.Require().Error(err)
}

// Test Runner
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
