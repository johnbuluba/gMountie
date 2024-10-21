package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	fullConf string
}

func (s *ConfigTestSuite) SetupTest() {
	s.fullConf = `
server:
  address: 0.0.0.0
  port: 8000
volumes:
  - name: test
    path: /tmp
`
}

// Tests
func (s *ConfigTestSuite) TestParse_Full_Server() {
	// Test.
	result, err := LoadConfigFromString(s.fullConf)

	// Verify.
	s.Require().NoError(err)
	s.Assert().Equal("test", result.Server.Address)
	s.Assert().Equal(uint(8000), result.Server.Port)
}

func (s *ConfigTestSuite) TestParse_Full_Volumes() {
	// Test.
	result, err := LoadConfigFromString(s.fullConf)

	// Verify.
	s.Require().NoError(err)
	s.Assert().Len(result.Volumes, 1)
	s.Assert().Equal("test", result.Volumes[0].Name)
	s.Assert().Equal("/tmp", result.Volumes[0].Path)
}

// Test Runner
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
