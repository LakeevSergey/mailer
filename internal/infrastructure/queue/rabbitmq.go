package queue

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/application"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
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

func (r *RabbitMQ[T]) Save(sendMail entity.SendMail) error {
	return errors.New("not implemented")
}

func (r *RabbitMQ[T]) Listen(ctx context.Context, worker func(context.Context, entity.SendMail) error) error {
	return errors.New("not implemented")
}
