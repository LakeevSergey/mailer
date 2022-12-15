package mailsender

import (
	"github.com/LakeevSergey/mailer/internal/common/dto"
)

type MailSender struct {
	savier RequestSavier
}

func NewMailSender(savier RequestSavier) *MailSender {
	return &MailSender{
		savier: savier,
	}
}

func (s *MailSender) Send(sendMail dto.SendMail) error {
	return s.savier.Save(sendMail)
}
