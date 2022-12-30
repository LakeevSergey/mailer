package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	SendFromEmail string `env:"SEND_FROM_EMAIL"`
	SendFromName  string `env:"SEND_FROM_NAME"`

	DBHost     string `env:"MYSQL_HOST"`
	DBPort     int    `env:"MYSQL_PORT"`
	DBName     string `env:"MYSQL_DATABASE"`
	DBUser     string `env:"MYSQL_USER"`
	DBPassword string `env:"MYSQL_PASSWORD"`

	ConsoleLoggerLevel int `env:"CONSOLE_LOGGER_LEVEL"`

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
