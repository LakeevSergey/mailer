package listner

import (
	"context"
	"errors"
)

type RabbitMQListner[T any] struct {
	decoder Decoder[T]
}

func NewRabbitMQListner[T any](decoder Decoder[T]) *RabbitMQListner[T] {
	return &RabbitMQListner[T]{
		decoder: decoder,
	}
}

func (l *RabbitMQListner[T]) Listen(ctx context.Context, worker func(T) error) error {
	return errors.New("not implemented")
}
