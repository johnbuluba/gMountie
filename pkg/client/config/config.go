package config

import (
	"gmountie/pkg/common/config"
	serverConfig "gmountie/pkg/server/config"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config is a struct that holds the configuration for the client
type Config struct {
	// Server is the server configuration
	Server *ServerConfig `validate:"required"`
	// Auth is the authentication configuration
	Auth serverConfig.AuthConfig `validate:"required"`
}

// LoadConfigFromString loads a Config from a string
func LoadConfigFromString(cfg string) (*Config, error) {
	return config.LoadConfigFromString(cfg, ParseConfig)
}

// ParseConfig parses a Config from a viper.Viper
func ParseConfig(v *viper.Viper) (*Config, error) {
	var result Config

	// Parse server config
	v.SetDefault("server", make(map[string]string))
	if cfg, err := NewServerConfig(v.Sub("server")); err == nil {
		result.Server = cfg
	} else {
		return nil, err
	}
	// Parse auth config
	v.SetDefault("auth", make(map[string]string))
	if cfg, err := NewAuthFromConfig(v.Sub("auth")); err == nil {
		result.Auth = cfg
	} else {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(result); err != nil {
		return nil, err
	}
	return &result, nil
}
