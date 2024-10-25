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
auth:
  type: none	
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
	s.Assert().Equal("0.0.0.0", result.Server.Address)
	s.Assert().Equal(uint(8000), result.Server.Port)
	s.Assert().Equal(AuthConfigTypeNone, result.Auth.GetType())
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

// TestParse_BasicAuthConfig
func (s *ConfigTestSuite) TestParse_BasicAuthConfig() {
	// Setup.
	conf := `
server:
  address: 0.0.0.0
  port: 8000

auth:
  type: basic
  users:
  - username: test
    password: test
volumes:
  - name: test
    path: /tmp
`
	// Test.
	result, err := LoadConfigFromString(conf)

	// Verify.
	s.Require().NoError(err)
	s.Assert().Equal(AuthConfigTypeBasic, result.Auth.GetType())
	s.Assert().Len(result.Auth.(*BasicAuthConfig).Users, 1)
	s.Assert().Equal("test", result.Auth.(*BasicAuthConfig).Users[0].Username)
	s.Assert().Equal("test", result.Auth.(*BasicAuthConfig).Users[0].Password)
}

// TestParse_BasicAuthConfig_Invalid
func (s *ConfigTestSuite) TestParse_BasicAuthConfig_Invalid() {
	// Setup.
	conf := `
server:
  address: 0.0.0.0
  port: 8000

auth:
  type: basic
  users:
  - username: nopassword
volumes:
  - name: test
    path: /tmp
`
	// Test.
	_, err := LoadConfigFromString(conf)

	// Verify.
	s.Require().Error(err)
}

// Test Runner
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
