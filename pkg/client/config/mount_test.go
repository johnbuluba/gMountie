package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MountConfigTestSuite struct {
	suite.Suite
}

// Test successful parsing of "single" mount configuration
func (s *MountConfigTestSuite) TestParse_SingleMount() {
	conf := `
mount:
  type: single
  path: /mnt/test
  volume: testvol
` + minimalServerConfig
	result, err := LoadConfigFromString(conf, "")
	s.Require().NoError(err)
	s.Assert().Equal(MountTypeSingle, result.Mount.GetType())

	singleMount, ok := result.Mount.(*SingleMountConfig)
	s.Require().True(ok)
	s.Assert().Equal("/mnt/test", singleMount.Path)
	s.Assert().Equal("testvol", singleMount.Volume)
}

// Test successful parsing of VFS mount configuration
func (s *MountConfigTestSuite) TestParse_VFSMount() {
	conf := `
mount:
  type: vfs
  path: /mnt/vfs
  mount_all: true
  volumes:
    - vol1
    - vol2
` + minimalServerConfig
	result, err := LoadConfigFromString(conf, "")
	s.Require().NoError(err)
	s.Assert().Equal(MountTypeVFS, result.Mount.GetType())

	vfsMount, ok := result.Mount.(*VFSMountConfig)
	s.Require().True(ok)
	s.Assert().Equal("/mnt/vfs", vfsMount.Path)
	s.Assert().True(vfsMount.MountAll)
	s.Assert().ElementsMatch([]string{"vol1", "vol2"}, vfsMount.Volumes)
}

// Test error cases for invalid mount type
func (s *MountConfigTestSuite) TestParse_InvalidMountType() {
	conf := `
mount:
  type: invalid
  path: /mnt/test
` + minimalServerConfig
	_, err := LoadConfigFromString(conf, "")
	s.Require().Error(err)
	s.Assert().Contains(err.Error(), "invalid mount type")
}

// Test error cases for missing required fields
func (s *MountConfigTestSuite) TestParse_MissingRequiredFields() {
	// Missing path
	conf := `
mount:
  type: single
  volume: testvol
` + minimalServerConfig
	_, err := LoadConfigFromString(conf, "")
	s.Require().Error(err)
}

// Test VFS mount with default values
func (s *MountConfigTestSuite) TestParse_VFSMountDefaults() {
	conf := `
mount:
  type: vfs
  path: /mnt/vfs
` + minimalServerConfig
	result, err := LoadConfigFromString(conf, "")
	s.Require().NoError(err)

	vfsMount, ok := result.Mount.(*VFSMountConfig)
	s.Require().True(ok)
	s.Assert().False(vfsMount.MountAll)
	s.Assert().Empty(vfsMount.Volumes)
}

func TestMountConfigTestSuite(t *testing.T) {
	suite.Run(t, new(MountConfigTestSuite))
}
