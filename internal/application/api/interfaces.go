package api

import (
	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type Decoder interface {
	Decode(data []byte) (entity.SendMail, error)
}

type MailSender interface {
	Send(sendMail entity.SendMail) error
}
