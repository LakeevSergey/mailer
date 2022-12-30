package mailsender

import "github.com/LakeevSergey/mailer/internal/domain/entity"

type MailSender struct {
	savier RequestSavier
}

func NewMailSender(savier RequestSavier) *MailSender {
	return &MailSender{
		savier: savier,
	}
}

func (s *MailSender) Send(sendMail entity.SendMail) error {
	return s.savier.Save(sendMail)
}
