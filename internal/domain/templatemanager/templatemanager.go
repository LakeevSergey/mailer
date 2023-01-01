package templatemanager

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type TemplateManager struct {
	storager TemplateStorager
}

func NewTemplateManager(storager TemplateStorager) *TemplateManager {
	return &TemplateManager{
		storager: storager,
	}
}

func (m *TemplateManager) Get(ctx context.Context, id int) (entity.Template, error) {
	return m.storager.Get(ctx, id)
}

func (m *TemplateManager) Search(ctx context.Context, dto dto.Search) (templates []entity.Template, total int, err error) {
	return m.storager.Search(ctx, dto)
}

func (m *TemplateManager) Add(ctx context.Context, dto dto.Add) (entity.Template, error) {
	return m.storager.Add(ctx, dto)
}

func (m *TemplateManager) Update(ctx context.Context, id int, dto dto.Update) (entity.Template, error) {
	return m.storager.Update(ctx, id, dto)
}

func (m *TemplateManager) Delete(ctx context.Context, id int) error {
	return m.storager.Delete(ctx, id)
}
