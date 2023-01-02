package templatemanager

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type TemplateManager interface {
	Get(ctx context.Context, id int64) (entity.Template, error)
	Search(ctx context.Context, dto dto.Search) (templates []entity.Template, total int, err error)
	Add(ctx context.Context, dto dto.Add) (entity.Template, error)
	Update(ctx context.Context, id int64, dto dto.Update) (entity.Template, error)
	Delete(ctx context.Context, id int64) error
}
