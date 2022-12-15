package domain

import (
	"github.com/LakeevSergey/mailer/internal/common/dto"
	"github.com/LakeevSergey/mailer/internal/consumer/domain/entity"
)

type Mailer struct {
	storager        TemplateStorager
	builder         MailBuilder
	sender          MailSender
	defaultSendFrom entity.SendFrom
}

func NewMailer(storager TemplateStorager, builder MailBuilder, sender MailSender, defaultSendFrom entity.SendFrom) *Mailer {
	return &Mailer{
		storager:        storager,
		builder:         builder,
		sender:          sender,
		defaultSendFrom: defaultSendFrom,
	}
}

func (m *Mailer) Send(sendMail dto.SendMail) error {
	template, err := m.storager.Get(sendMail.Code)
	if err != nil {
		return err
	}

	body, title, err := m.builder.Build(template, sendMail.Params)
	if err != nil {
		return err
	}

	sendFrom := m.defaultSendFrom

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

	return m.sender.Send(mail)
}
