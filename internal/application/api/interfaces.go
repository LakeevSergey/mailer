package api

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/application/dto"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	templatemanagerdto "github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type Decoder interface {
	Decode(data []byte) (dto.SendMail, error)
}

type MailSender interface {
	Send(sendMail entity.SendMail) error
}

type TemplateManager interface {
	Get(ctx context.Context, id int) (entity.Template, error)
	Search(ctx context.Context, dto templatemanagerdto.Search) (templates []entity.Template, total int, err error)
	Add(ctx context.Context, dto templatemanagerdto.Add) (entity.Template, error)
	Update(ctx context.Context, id int, dto templatemanagerdto.Update) (entity.Template, error)
	Delete(ctx context.Context, id int) error
}
