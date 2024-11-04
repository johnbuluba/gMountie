package config

import (
	"os"
	"path/filepath"
)

const (
	DefaultConfigFileName = "config.yaml"
	DefaultConfigDirName  = "gmountie"
)

// GetDefaultConfigPath returns the default config file path based on the OS
func GetDefaultConfigPath() string {
	configDir := GetDefaultConfigDir()
	return filepath.Join(configDir, DefaultConfigFileName)
}

// GetDefaultConfigDir returns the default config directory for Linux
func GetDefaultConfigDir() string {
	// Linux/Unix: $XDG_CONFIG_HOME/gmountie or ~/.config/gmountie
	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	homeDir, _ := os.UserHomeDir()
	if xdgConfig != "" {
		return filepath.Join(xdgConfig, DefaultConfigDirName)
	}
	return filepath.Join(homeDir, ".config", DefaultConfigDirName)
}

// EnsureConfigDir creates the config directory if it doesn't exist
func EnsureConfigDir() error {
	configDir := GetDefaultConfigDir()
	return os.MkdirAll(configDir, 0755)
}

// WriteDefaultConfig writes the default config to the default location
func WriteDefaultConfig() error {
	configPath := GetDefaultConfigPath()
	return os.WriteFile(configPath, []byte(DefaultConfig), 0644)
}

// DefaultConfig is the default configuration
const DefaultConfig = `server:
  address: 127.0.0.1
  port: 9449
  metrics: true

auth:
  type: basic
  users:
    - username: admin
      password: admin
`
