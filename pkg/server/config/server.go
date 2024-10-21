package config

import (
	"github.com/spf13/viper"
)

const (
	// DefaultAddress is the default address that the server will listen on
	DefaultAddress = "0.0.0.0"
	// DefaultPort is the default port that the server will listen on
	DefaultPort = 9449
)

// ServerConfig is a struct that holds the configuration for the server
type ServerConfig struct {
	// Address is the address that the server will listen on
	Address string `validate:"required,ip"`
	// Port is the port that the server will listen on
	Port uint `validate:"required"`
}

// NewServerConfig creates a new ServerConfig with defaults
func NewServerConfig(v *viper.Viper) *ServerConfig {
	v.SetDefault("address", DefaultAddress)
	v.SetDefault("port", DefaultPort)

	return &ServerConfig{
		Address: v.GetString("address"),
		Port:    v.GetUint("port"),
	}
}
