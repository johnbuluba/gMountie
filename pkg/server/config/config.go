package config

import (
	"fmt"
	"gmountie/pkg/common/config"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const (
	EnvironmentPrefix = "GMOUNTIE"
)

// Config is a struct that holds the configuration for the server
type Config struct {
	// Server is the server configuration
	Server *ServerConfig `validate:"required"`

	// Auth is the auth configuration
	Auth AuthConfig `validate:"required"`

	// Volumes is the volume configuration
	Volumes []*VolumeConfig `validate:"required,dive"`
}

func LoadConfigFromString(cfg string) (*Config, error) {
	return config.LoadConfigFromString(cfg, ParseConfig)
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var result Config
	// Enable environment variables.
	v.SetEnvPrefix(EnvironmentPrefix)
	v.AutomaticEnv()

	// Parse the server configuration.
	v.SetDefault("server", make(map[string]string))
	result.Server = NewServerConfig(v.Sub("server"))

	// Parse the auth configuration.
	auth, err := NewFromConfig(v.Sub("auth"))
	if err != nil {
		return nil, err
	}
	result.Auth = auth

	// Parse the volume configuration.
	volumes := make([]*VolumeConfig, 0)
	for sub, i := v.Sub("volumes.0"), 0; sub != nil; sub = v.Sub(fmt.Sprintf("volumes.%d", i)) {
		volumes = append(volumes, NewVolumeConfig(sub))
		i++
	}
	result.Volumes = volumes

	// Validate.
	validate := validator.New()
	err = validate.Struct(result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
