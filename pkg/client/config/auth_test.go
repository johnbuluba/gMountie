package config

import (
	"testing"

	serverConfig "gmountie/pkg/server/config"

	"github.com/stretchr/testify/suite"
)

type AuthConfigTestSuite struct {
	suite.Suite
}

// Test successful parsing of "none" auth configuration
func (s *AuthConfigTestSuite) TestParse_NoneAuth() {
	conf := `
server:
  address: 127.0.0.1
auth:
  type: none
`
	result, err := LoadConfigFromString(conf)
	s.Require().NoError(err)
	s.Assert().Equal(serverConfig.AuthConfigTypeNone, result.Auth.GetType())
}

// Test successful parsing of basic auth configuration
func (s *AuthConfigTestSuite) TestParse_BasicAuth() {
	conf := `
server:
  address: 127.0.0.1
auth:
  type: basic
  username: testuser
  password: testpass
`
	result, err := LoadConfigFromString(conf)
	s.Require().NoError(err)
	s.Assert().Equal(serverConfig.AuthConfigTypeBasic, result.Auth.GetType())

	basicAuth, ok := result.Auth.(*BasicAuthConfig)
	s.Require().True(ok)
	s.Assert().Equal("testuser", basicAuth.Username)
	s.Assert().Equal("testpass", basicAuth.Password)
	test, err := result.String()
	s.Require().NoError(err)
	s.Assert().Contains(test, "testuser")
}

// Test error cases for invalid auth type
func (s *AuthConfigTestSuite) TestParse_InvalidAuthType() {
	conf := `
server:
  address: 127.0.0.1
auth:
  type: invalid
`
	_, err := LoadConfigFromString(conf)
	s.Require().Error(err)
	s.Assert().Contains(err.Error(), "invalid auth type")
}

// Test error cases for missing required fields in basic auth
func (s *AuthConfigTestSuite) TestParse_MissingBasicAuthFields() {
	// Missing username
	conf1 := `
server:
  address: 127.0.0.1
auth:
  type: basic
  password: testpass
`
	_, err := LoadConfigFromString(conf1)
	s.Require().Error(err)

	// Missing password
	conf2 := `
server:
  address: 127.0.0.1
auth:
  type: basic
  username: testuser
`
	_, err = LoadConfigFromString(conf2)
	s.Require().Error(err)
}

// Test validation of username/password requirements
func (s *AuthConfigTestSuite) TestParse_EmptyCredentials() {
	conf := `
auth:
  type: basic
  username: ""
  password: ""
`
	_, err := LoadConfigFromString(conf + minimalServerConfig)
	s.Require().Error(err)
}

// Test that GetType() returns correct auth type
func (s *AuthConfigTestSuite) TestGetType() {
	basicAuth := &BasicAuthConfig{
		BasicAuthConfigUser: serverConfig.BasicAuthConfigUser{
			Username: "test",
			Password: "test",
		},
	}
	s.Assert().Equal(serverConfig.AuthConfigTypeBasic, basicAuth.GetType())
}

// Test integration with the main config parser
func (s *AuthConfigTestSuite) TestIntegration_WithServerConfig() {
	conf := `
server:
  address: 127.0.0.1
  port: 9449
auth:
  type: basic
  username: testuser
  password: testpass
`
	result, err := LoadConfigFromString(conf)
	s.Require().NoError(err)
	s.Assert().NotNil(result.Server)
	s.Assert().NotNil(result.Auth)
	s.Assert().Equal(serverConfig.AuthConfigTypeBasic, result.Auth.GetType())
}

var minimalServerConfig = `
server:
  address: 127.0.0.1
auth:
  type: none
`

func TestAuthConfigTestSuite(t *testing.T) {
	suite.Run(t, new(AuthConfigTestSuite))
}
