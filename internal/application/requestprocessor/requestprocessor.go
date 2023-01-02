package requestprocessor

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type SendMailRequestProcessor struct {
	storager        TemplateStorager
	builder         MailBuilder
	sender          MailSender
	defaultSendFrom entity.SendFrom
}

func NewSendMailRequestProcessor(storager TemplateStorager, builder MailBuilder, sender MailSender, defaultSendFrom entity.SendFrom) *SendMailRequestProcessor {
	return &SendMailRequestProcessor{
		storager:        storager,
		builder:         builder,
		sender:          sender,
		defaultSendFrom: defaultSendFrom,
	}
}

func (p *SendMailRequestProcessor) Process(ctx context.Context, sendMail entity.SendMail) error {
	template, err := p.storager.GetByCode(ctx, sendMail.Code)
	if err != nil {
		return err
	}

	body, title, err := p.builder.Build(template, sendMail.Params)
	if err != nil {
		return err
	}

	sendFrom := p.defaultSendFrom

	if sendMail.SendFrom != nil {
		sendFrom = entity.SendFrom{
			Email: sendMail.SendFrom.Email,
			Name:  sendMail.SendFrom.Name,
		}
	}

	mail := entity.Mail{
		SendTo:   sendMail.SendTo,
		SendFrom: sendFrom,
		Title:    title,
		Body:     body,
	}

	return p.sender.Send(mail)
}
