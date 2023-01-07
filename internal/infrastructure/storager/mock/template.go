package mock

import (
	"context"
	"errors"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

var ActiveTemplate = entity.Template{
	Id:     1,
	Active: true,
	Code:   "active_template",
	Name:   "Active template",
	Body:   "Active template body, {{ name }}",
	Title:  "Active template title, {{ name }}",
}

var NotActiveTemplate = entity.Template{
	Id:     2,
	Active: false,
	Code:   "not_active_template",
	Name:   "Not active template",
	Body:   "Not active template body, {{ name }}",
	Title:  "Not active template title, {{ name }}",
}

type MockTemplateStorager struct {
	mock.Mock
}

func NewMockStorager(ctx context.Context) *MockTemplateStorager {
	mockStorager := &MockTemplateStorager{}

	mockStorager.On("GetByCode", ctx, "active_template").Return(ActiveTemplate, nil)
	mockStorager.On("GetByCode", ctx, "not_active_template").Return(NotActiveTemplate, nil)
	mockStorager.On("GetByCode", ctx, "not_existent_template").Return(entity.Template{}, errors.New("entity not found"))

	mockStorager.On("Get", ctx, 1).Return(ActiveTemplate, nil)
	mockStorager.On("Get", ctx, 2).Return(NotActiveTemplate, nil)
	mockStorager.On("Get", ctx, 3).Return(entity.Template{}, errors.New("entity not found"))

	return mockStorager
}

func (s *MockTemplateStorager) GetByCode(ctx context.Context, code string) (entity.Template, error) {
	args := s.Called(ctx, code)
	return args.Get(0).(entity.Template), args.Error(1)
}
