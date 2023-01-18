package templatemanager

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager"
	storagerdto "github.com/LakeevSergey/mailer/internal/domain/storager/dto"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type TemplateManager struct {
	storager storager.TemplateStorager
}

func NewTemplateManager(storager storager.TemplateStorager) *TemplateManager {
	return &TemplateManager{
		storager: storager,
	}
}

func (m *TemplateManager) Get(ctx context.Context, id int64) (entity.Template, error) {
	template, err := m.storager.Get(ctx, id)
	if errors.Is(err, storager.ErrorEntityNotFound) {
		return entity.Template{}, domain.ErrorTemplateNotFound
	} else if err != nil {
		return entity.Template{}, err
	}

	return template, nil
}

func (m *TemplateManager) Search(ctx context.Context, dto dto.SearchTemplate) (templates []entity.Template, total int, err error) {
	templates, total, err = m.storager.Search(ctx, storagerdto.SearchTemplate{
		Limit:  dto.Limit,
		Offset: dto.Offset,
	})
	if err != nil {
		return []entity.Template{}, 0, err
	}

	return templates, total, err
}

func (m *TemplateManager) Add(ctx context.Context, dto dto.AddTemplate) (entity.Template, error) {
	template, err := m.storager.Add(ctx, storagerdto.AddTemplate{
		Active: dto.Active,
		Code:   dto.Code,
		Name:   dto.Name,
		Body:   dto.Body,
		Title:  dto.Title,
	})
	if errors.Is(err, storager.ErrorDuplicate) {
		return entity.Template{}, domain.ErrorTemplateCodeDuplicate
	} else if err != nil {
		return entity.Template{}, err
	}

	return template, nil
}

func (m *TemplateManager) Update(ctx context.Context, id int64, dto dto.UpdateTemplate) (entity.Template, error) {
	template, err := m.storager.Update(ctx, id, storagerdto.UpdateTemplate{
		Active: dto.Active,
		Code:   dto.Code,
		Name:   dto.Name,
		Body:   dto.Body,
		Title:  dto.Title,
	})
	if errors.Is(err, storager.ErrorDuplicate) {
		return entity.Template{}, domain.ErrorTemplateCodeDuplicate
	} else if errors.Is(err, storager.ErrorEntityNotFound) {
		return entity.Template{}, domain.ErrorTemplateNotFound
	} else if err != nil {
		return entity.Template{}, err
	}

	return template, nil
}

func (m *TemplateManager) Delete(ctx context.Context, id int64) error {
	err := m.storager.Delete(ctx, id)
	if errors.Is(err, storager.ErrorEntityNotFound) {
		return domain.ErrorTemplateNotFound
	} else if err != nil {
		return err
	}

	return nil
}
