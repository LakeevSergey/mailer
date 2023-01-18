package requestprocessor

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/attachmentmanager"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type SendMailRequestProcessor struct {
	storager          TemplateStorager
	builder           MailBuilder
	sender            MailSender
	attachmentManager attachmentmanager.AttachmentManager
	defaultSendFrom   entity.SendFrom
}

func NewSendMailRequestProcessor(storager TemplateStorager, builder MailBuilder, sender MailSender, attachmentManager attachmentmanager.AttachmentManager, defaultSendFrom entity.SendFrom) *SendMailRequestProcessor {
	return &SendMailRequestProcessor{
		storager:          storager,
		builder:           builder,
		sender:            sender,
		attachmentManager: attachmentManager,
		defaultSendFrom:   defaultSendFrom,
	}
}

func (p *SendMailRequestProcessor) Process(ctx context.Context, sendMail entity.SendMail) error {
	template, err := p.storager.GetByCode(ctx, sendMail.Code)
	if err != nil {
		return err
	}

	if !template.Active {
		return ErrorTemplateDeactivated
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

	attachments := make([]entity.File, 0, len(sendMail.Attachments))
	for _, fileId := range sendMail.Attachments {
		file, err := p.attachmentManager.Get(ctx, fileId)
		if err != nil {
			return err
		}
		file.Data.Close()

		attachments = append(attachments, file)
	}

	mail := entity.Mail{
		SendTo:      sendMail.SendTo,
		SendFrom:    sendFrom,
		Title:       title,
		Body:        body,
		Attachments: attachments,
	}

	return p.sender.Send(mail)
}
