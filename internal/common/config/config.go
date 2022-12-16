package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	SendFromEmail string `env:"SEND_FROM_EMAIL"`
	SendFromName  string `env:"SEND_FROM_NAME"`

	ApiPort int `env:"API_PORT"`
}

func New() (Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
