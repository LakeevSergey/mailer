package mailsender

import "github.com/LakeevSergey/mailer/internal/common/dto"

type RequestSavier interface {
	Save(sendMail dto.SendMail) error
}
