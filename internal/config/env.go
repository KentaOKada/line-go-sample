package config

import (
	"github.com/kelseyhightower/envconfig"
)

func LoadFromEnv() (*Config, error) {
	var conf Config

	if err := envconfig.Process("", &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
