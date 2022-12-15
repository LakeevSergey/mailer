package sender

import (
	"errors"

	"github.com/LakeevSergey/mailer/internal/consumer/domain/entity"
)

type SMTPSender struct {
}

func NewSMTPSender() *SMTPSender {
	return &SMTPSender{}
}

func (s SMTPSender) Send(mail entity.Mail) error {
	return errors.New("not implemented")
}
