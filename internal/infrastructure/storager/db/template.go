package db

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
)

type DBTemplateStorager struct {
}

func NewDBTemplateStorager() *DBTemplateStorager {
	return &DBTemplateStorager{}
}

func (s *DBTemplateStorager) GetByCode(ctx context.Context, code string) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}

func (s *DBTemplateStorager) Get(ctx context.Context, id int) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}

func (s *DBTemplateStorager) Search(ctx context.Context, dto dto.Search) (templates []entity.Template, total int, err error) {
	return []entity.Template{}, 0, errors.New("not implemented")
}

func (s *DBTemplateStorager) Add(ctx context.Context, dto dto.Add) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}

func (s *DBTemplateStorager) Update(ctx context.Context, id int, dto dto.Update) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}

func (s *DBTemplateStorager) Delete(ctx context.Context, id int) error {
	return errors.New("not implemented")
}
