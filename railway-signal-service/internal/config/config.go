package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PG_ADDR     string `env:"PG_ADDR" default:"localhost:5432"`
	PG_USER     string `env:"PG_USER" default:"user"`
	PG_PASSWORD string `env:"PG_PASSWORD" default:"password"`
	PG_APP_NAME string `env:"PG_APP_NAME" default:"railway"`
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
