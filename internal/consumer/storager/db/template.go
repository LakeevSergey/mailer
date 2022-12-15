package db

import (
	"errors"

	"github.com/LakeevSergey/mailer/internal/consumer/domain/entity"
)

type DBTemplateStorager struct {
}

func NewDBTemplateStorager() *DBTemplateStorager {
	return &DBTemplateStorager{}
}

func (s *DBTemplateStorager) Get(code string) (entity.Template, error) {
	return entity.Template{}, errors.New("not implemented")
}
