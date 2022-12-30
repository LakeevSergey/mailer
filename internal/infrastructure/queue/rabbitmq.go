package queue

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/application"
)

type Config struct {
	User          string
	Password      string
	Host          string
	Port          int
	Queue         string
	ExchangeInput string
	ExchangeDLX   string
	QueueDLX      string
	RetryDelay    int
}

type RabbitMQ[T any] struct {
	config Config
	coder  Coder[T]
	logger application.Logger
}

func NewRabbitMQ[T any](config Config, coder Coder[T], logger application.Logger) *RabbitMQ[T] {
	return &RabbitMQ[T]{
		config: config,
		coder:  coder,
		logger: logger,
	}
}

func (r *RabbitMQ[T]) Save(message T) error {
	return errors.New("not implemented")
}

func (r *RabbitMQ[T]) Listen(ctx context.Context, worker func(context.Context, T) error) error {
	return errors.New("not implemented")
}
