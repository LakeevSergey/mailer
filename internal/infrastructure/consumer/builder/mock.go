package builder

import (
	"github.com/stretchr/testify/mock"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	storagermock "github.com/LakeevSergey/mailer/internal/infrastructure/storager/mock"
)

type MockBuilder struct {
	mock.Mock
}

func NewMockBuilder() *MockBuilder {
	mockBuilder := &MockBuilder{}

	mockBuilder.On("Build", storagermock.ActiveTemplate, map[string]string{"name": "Name"}).Return(
		"Active template body, Name",
		"Active template title, Name",
		nil,
	)

	mockBuilder.On("Build", storagermock.NotActiveTemplate, map[string]string{"name": "Name"}).Return(
		"Not active template body, Name",
		"Not active template title, Name",
		nil,
	)

	return mockBuilder
}

func (b *MockBuilder) Build(template entity.Template, params map[string]string) (body string, title string, err error) {
	args := b.Called(template, params)

	return args.String(0), args.String(1), args.Error(2)
}
