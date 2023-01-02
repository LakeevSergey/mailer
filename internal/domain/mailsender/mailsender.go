package mailsender

import "github.com/LakeevSergey/mailer/internal/domain/entity"

type MailSender interface {
	Send(sendMail entity.SendMail) error
}
