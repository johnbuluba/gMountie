package config

import (
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

// Test Runner
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
