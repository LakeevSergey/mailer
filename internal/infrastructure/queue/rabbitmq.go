package queue

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/application"
)

type RabbitMQ[T any] struct {
	coder  Coder[T]
	logger application.Logger
}

func NewRabbitMQ[T any](coder Coder[T], logger application.Logger) *RabbitMQ[T] {
	return &RabbitMQ[T]{
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
