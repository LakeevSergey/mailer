package db

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
)

type DBTemplateStorager struct {
}

func NewDBTemplateStorager() *DBTemplateStorager {
	return &DBTemplateStorager{}
}

func (s *DBTemplateStorager) GetByCode(ctx context.Context, code string) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}
