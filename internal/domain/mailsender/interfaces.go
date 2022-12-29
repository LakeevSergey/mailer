package mailsender

import "github.com/LakeevSergey/mailer/internal/domain/entity"

type RequestSavier interface {
	Save(sendMail entity.SendMail) error
}
