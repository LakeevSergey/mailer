package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	SendFromEmail string `env:"SEND_FROM_EMAIL"`
	SendFromName  string `env:"SEND_FROM_NAME"`

	SMTPHost     string `env:"SMTP_HOST"`
	SMTPPort     int    `env:"SMTP_PORT"`
	SMTPUser     string `env:"SMTP_USER"`
	SMTPPassword string `env:"SMTP_PASSWORD"`

	DBHost     string `env:"MYSQL_HOST"`
	DBPort     int    `env:"MYSQL_PORT"`
	DBName     string `env:"MYSQL_DATABASE"`
	DBUser     string `env:"MYSQL_USER"`
	DBPassword string `env:"MYSQL_PASSWORD"`

	RBMQHost          string `env:"RBMQ_HOST"`
	RBMQPort          int    `env:"RBMQ_PORT"`
	RBMQUser          string `env:"RBMQ_USER"`
	RBMQPassword      string `env:"RBMQ_PASSWORD"`
	RBMQQueue         string `env:"RBMQ_QUEUE"`
	RBMQQueueDLX      string `env:"RBMQ_QUEUE_DLX"`
	RBMQExchangeDLX   string `env:"RBMQ_EXCHANGE_DLX"`
	RBMQExchangeInput string `env:"RBMQ_EXCHANGE_INPUT"`

	RetryCount int `env:"RETRY_COUNT"`
	RetryDelay int `env:"RETRY_DELAY"`

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
