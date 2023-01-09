package mock

import (
	"context"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/domain/storager"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
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

var AddSuccess = dto.Add{
	Active: true,
	Code:   "new_template",
	Name:   "New template",
	Body:   "New active template body, {{ name }}",
	Title:  "New active template title, {{ name }}",
}

var AddDublicateError = dto.Add{
	Active: true,
	Code:   "active_template",
	Name:   "New template",
	Body:   "New active template body, {{ name }}",
	Title:  "New active template title, {{ name }}",
}

var UpdateSuccess = dto.Update{
	Active: true,
	Code:   "new_template",
	Name:   "New template",
	Body:   "New active template body, {{ name }}",
	Title:  "New active template title, {{ name }}",
}

var UpdateDublicateError = dto.Update{
	Active: true,
	Code:   "active_template",
	Name:   "New template",
	Body:   "New active template body, {{ name }}",
	Title:  "New active template title, {{ name }}",
}

type MockTemplateStorager struct {
	mock.Mock
}

func NewMockStorager(ctx context.Context) *MockTemplateStorager {
	mockStorager := &MockTemplateStorager{}

	mockStorager.On("GetByCode", ctx, "active_template").Return(ActiveTemplate, nil)
	mockStorager.On("GetByCode", ctx, "not_active_template").Return(NotActiveTemplate, nil)
	mockStorager.On("GetByCode", ctx, "not_existent_template").Return(entity.Template{}, storager.ErrorEntityNotFound)

	mockStorager.On("Get", ctx, int64(1)).Return(ActiveTemplate, nil)
	mockStorager.On("Get", ctx, int64(2)).Return(NotActiveTemplate, nil)
	mockStorager.On("Get", ctx, int64(3)).Return(entity.Template{}, storager.ErrorEntityNotFound)

	mockStorager.On("Search", ctx, dto.Search{Limit: 10, Offset: 0}).Return(
		[]entity.Template{
			ActiveTemplate,
			NotActiveTemplate,
		},
		2,
		nil,
	)
	mockStorager.On("Search", ctx, dto.Search{Limit: 10, Offset: 10}).Return(
		[]entity.Template{},
		2,
		nil,
	)

	mockStorager.On("Search", ctx, dto.Search{Limit: 10, Offset: 10}).Return(
		[]entity.Template{},
		2,
		nil,
	)

	mockStorager.On("Add", ctx, AddSuccess).Return(
		entity.Template{
			Id:     3,
			Active: AddSuccess.Active,
			Code:   AddSuccess.Code,
			Name:   AddSuccess.Name,
			Body:   AddSuccess.Body,
			Title:  AddSuccess.Title,
		},
		nil,
	)

	mockStorager.On("Add", ctx, AddDublicateError).Return(
		entity.Template{},
		storager.ErrorDuplicate,
	)

	mockStorager.On("Update", ctx, int64(2), UpdateSuccess).Return(
		entity.Template{
			Id:     2,
			Active: UpdateSuccess.Active,
			Code:   UpdateSuccess.Code,
			Name:   UpdateSuccess.Name,
			Body:   UpdateSuccess.Body,
			Title:  UpdateSuccess.Title,
		},
		nil,
	)

	mockStorager.On("Update", ctx, int64(3), UpdateSuccess).Return(
		entity.Template{},
		storager.ErrorEntityNotFound,
	)

	mockStorager.On("Update", ctx, int64(2), UpdateDublicateError).Return(
		entity.Template{},
		storager.ErrorDuplicate,
	)

	mockStorager.On("Delete", ctx, int64(2), UpdateDublicateError).Return(
		nil,
	)

	mockStorager.On("Delete", ctx, int64(3), UpdateSuccess).Return(
		storager.ErrorEntityNotFound,
	)

	return mockStorager
}

func (s *MockTemplateStorager) GetByCode(ctx context.Context, code string) (entity.Template, error) {
	args := s.Called(ctx, code)
	return args.Get(0).(entity.Template), args.Error(1)
}

func (s *MockTemplateStorager) Get(ctx context.Context, id int64) (entity.Template, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(entity.Template), args.Error(1)
}

func (s *MockTemplateStorager) Search(ctx context.Context, dto dto.Search) (templates []entity.Template, total int, err error) {
	args := s.Called(ctx, dto)
	return args.Get(0).([]entity.Template), args.Int(1), args.Error(2)
}

func (s *MockTemplateStorager) Add(ctx context.Context, dto dto.Add) (entity.Template, error) {
	args := s.Called(ctx, dto)
	return args.Get(0).(entity.Template), args.Error(1)
}

func (s *MockTemplateStorager) Update(ctx context.Context, id int64, dto dto.Update) (entity.Template, error) {
	args := s.Called(ctx, id, dto)
	return args.Get(0).(entity.Template), args.Error(1)
}

func (s *MockTemplateStorager) Delete(ctx context.Context, id int64) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}
