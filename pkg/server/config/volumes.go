package config

import "github.com/spf13/viper"

// VolumeConfig is a struct that holds the configuration for the volume
type VolumeConfig struct {
	// Name is the name of the volume
	Name string `validate:"required"`
	// Path is the path of the volume
	Path string `validate:"required"`
}

// NewVolumeConfig creates a new VolumeConfig with defaults
func NewVolumeConfig(v *viper.Viper) *VolumeConfig {
	return &VolumeConfig{
		Name: v.GetString("name"),
		Path: v.GetString("path"),
	}
}
