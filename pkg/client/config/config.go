package config

import (
	"gmountie/pkg/common/config"
	serverConfig "gmountie/pkg/server/config"
	"gmountie/pkg/utils/log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Config is a struct that holds the configuration for the client
type Config struct {
	configPath string
	// Server is the server configuration
	Server *ServerConfig `validate:"required"`
	// Auth is the authentication configuration
	Auth serverConfig.AuthConfig `validate:"required"`
	// Mount is the mount configuration
	Mount MountConfig `yaml:"mount,omitempty"`
}

// String returns the string representation of the Config
func (c *Config) String() (string, error) {
	out, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Save saves the Config to a file
func (c *Config) Save() error {
	content, err := c.String()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(c.configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		log.Log.Error("error closing file", zap.Error(err))
	}()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

// LoadConfigFromString loads a Config from a string
func LoadConfigFromString(cfg string, path string) (*Config, error) {
	c, err := config.LoadConfigFromString(cfg, ParseConfig)
	if err != nil {
		return nil, err
	}
	c.configPath = path
	return c, nil
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
	// Parse mount config
	mount := v.Sub("mount")
	if mount != nil {
		if cfg, err := NewMountConfig(v.Sub("mount")); err == nil {
			result.Mount = cfg
		} else {
			return nil, err
		}
	}

	validate := validator.New()
	if err := validate.Struct(result); err != nil {
		return nil, err
	}
	return &result, nil
}
