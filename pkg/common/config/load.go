package config

import (
	"bytes"

	"github.com/spf13/viper"
)

const (
	EnvironmentPrefix = "GMOUNTIE"
)

type Parser[T any] func(v *viper.Viper) (*T, error)

func LoadConfigFromString[T any](cfg string, parser Parser[T]) (*T, error) {
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBufferString(cfg))
	if err != nil {
		return nil, err
	}
	return parser(viper.GetViper())
}
