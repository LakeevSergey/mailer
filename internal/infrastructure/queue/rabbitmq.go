package queue

import (
	"context"
	"errors"
	"fmt"

	"github.com/LakeevSergey/mailer/internal/application"
	"github.com/streadway/amqp"
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
	connect *amqp.Connection
	queue   *amqp.Queue
	channel *amqp.Channel
	coder   Coder[T]
	logger  application.Logger
}

func NewRabbitMQ[T any](cfg Config, coder Coder[T], logger application.Logger) (*RabbitMQ[T], error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	// exchange для отправки сообщений в основную очередь
	err = channel.ExchangeDeclare(
		cfg.ExchangeInput,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	// exchange для отправки сообщений в очередь повторной обработки
	err = channel.ExchangeDeclare(
		cfg.ExchangeDLX,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	// очередь сообщений для повторной обработки
	queueDlx, err := channel.QueueDeclare(
		cfg.QueueDLX,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-message-ttl":          cfg.RetryDelay,
			"x-dead-letter-exchange": cfg.ExchangeInput,
		},
	)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind(
		queueDlx.Name,
		"",
		cfg.ExchangeDLX,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// основная очередь
	queue, err := channel.QueueDeclare(
		cfg.Queue,
		true,
		false,
		false,
		false,
		amqp.Table{"x-dead-letter-exchange": cfg.ExchangeDLX},
	)
	if err != nil {
		return nil, err
	}

	err = channel.QueueBind(
		queue.Name,
		"",
		cfg.ExchangeInput,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ[T]{
		connect: conn,
		queue:   &queue,
		channel: channel,
		coder:   coder,
		logger:  logger,
	}, nil
}

func (r *RabbitMQ[T]) Close() {
	r.channel.Close()
	r.connect.Close()
}

func (r *RabbitMQ[T]) Save(message T) error {
	return errors.New("not implemented")
}

func (r *RabbitMQ[T]) Listen(ctx context.Context, worker func(context.Context, T) error) error {
	return errors.New("not implemented")
}
