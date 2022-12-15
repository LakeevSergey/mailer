package api

import "github.com/LakeevSergey/mailer/internal/common/dto"

type Decoder interface {
	Decode(data []byte) (dto.SendMail, error)
}

type MailSender interface {
	Send(sendMail dto.SendMail) error
}
