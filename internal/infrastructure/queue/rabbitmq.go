package queue

import (
	"context"
	"fmt"

	"github.com/LakeevSergey/mailer/internal/infrastructure"
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
	RetryCount    int
}

type RabbitMQ[T any] struct {
	connect    *amqp.Connection
	queue      *amqp.Queue
	channel    *amqp.Channel
	encoder    Encoder[T]
	logger     infrastructure.Logger
	retryCount int
}

func NewRabbitMQ[T any](cfg Config, encoder Encoder[T], logger infrastructure.Logger) (*RabbitMQ[T], error) {
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
		connect:    conn,
		queue:      &queue,
		channel:    channel,
		encoder:    encoder,
		logger:     logger,
		retryCount: cfg.RetryCount,
	}, nil
}

func (r *RabbitMQ[T]) Close() {
	r.channel.Close()
	r.connect.Close()
}

func (r *RabbitMQ[T]) Save(message T) error {
	body, err := r.encoder.Encode(message)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: r.encoder.ContentType(),
			Body:        body,
		},
	)
}

func (r *RabbitMQ[T]) Listen(ctx context.Context, worker func(context.Context, T) error) error {
	msgs, err := r.channel.Consume(
		r.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for {
		select {
		case delivery := <-msgs:
			err := func() error {
				data, err := r.encoder.Decode(delivery.Body)
				if err != nil {
					return err
				}
				return worker(ctx, data)
			}()

			if err != nil {
				var tryNumber int64 = 0

				xDeaths, ok := delivery.Headers["x-death"].([]interface{})
				if ok {
					var xDeath amqp.Table
					xDeath, ok = xDeaths[0].(amqp.Table)
					if ok {
						tryNumber, ok = xDeath["count"].(int64)
					}
				}

				if r.retryCount > 1 && (!ok || tryNumber < int64(r.retryCount)) {
					r.logger.WarnErr(fmt.Errorf("message process error, %d retries left: %w", int64(r.retryCount)-tryNumber, err))
					delivery.Nack(false, false)
				} else {
					r.logger.ErrorErr(fmt.Errorf("message process error, no retries left: %w", err))
					delivery.Ack(false)
				}
			} else {
				r.logger.Info("Mail sended!")
				delivery.Ack(false)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
