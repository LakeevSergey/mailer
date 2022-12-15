package requestsavier

import (
	"errors"

	"github.com/LakeevSergey/mailer/internal/common/dto"
)

type RabbitMQRequestSavier[T any] struct {
	encoder Encoder[T]
}

func NewRabbitMQRequestSavier[T any](encoder Encoder[T]) *RabbitMQRequestSavier[T] {
	return &RabbitMQRequestSavier[T]{
		encoder: encoder,
	}
}

func (s *RabbitMQRequestSavier[T]) Save(sendMail dto.SendMail) error {
	return errors.New("not implemented")
}
