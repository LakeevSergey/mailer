package domain

import "github.com/LakeevSergey/mailer/internal/consumer/domain/entity"

type TemplateStorager interface {
	Get(code string) (entity.Template, error)
}

type MailBuilder interface {
	Build(template entity.Template, params map[string]string) (body string, title string, err error)
}

type MailSender interface {
	Send(mail entity.Mail) error
}
