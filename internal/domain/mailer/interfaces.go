package mailer

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type TemplateStorager interface {
	GetByCode(ctx context.Context, code string) (entity.Template, error)
}

type MailBuilder interface {
	Build(template entity.Template, params map[string]string) (body string, title string, err error)
}

type MailSender interface {
	Send(mail entity.Mail) error
}
