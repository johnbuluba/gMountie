package config

import (
	"gmountie/pkg/server/config"

	"github.com/spf13/viper"
)

// ServerConfig is a struct that holds the configuration for the server
type ServerConfig struct {
	// Address is the address that the server will listen on
	Address string `validate:"required,ip"`
	// Port is the port that the server will listen on
	Port uint `validate:"required,gte=1,lte=65535"`
	// TLS
	TLS bool
}

// NewServerConfig creates a new ServerConfig with defaults
func NewServerConfig(v *viper.Viper) (*ServerConfig, error) {
	v.SetDefault("port", config.DefaultPort)
	v.SetDefault("tls", false)

	cfg := &ServerConfig{}
	if err := v.UnmarshalExact(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
