package sender

import (
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

var SuccessMail = entity.Mail{
	SendTo: []string{"test@test.test"},
	SendFrom: entity.SendFrom{
		Name:  "test",
		Email: "test@test.test",
	},
	Title: "Active template title, Name",
	Body:  "Active template body, Name",
}

var ErrorMail = entity.Mail{
	SendTo: []string{"test@test.test"},
	SendFrom: entity.SendFrom{
		Name:  "test",
		Email: "test@test.test",
	},
	Title: "Error",
	Body:  "Error",
}

type MockSender struct {
	mock.Mock
}

func NewMockSender() *MockSender {
	mockBuilder := &MockSender{}

	mockBuilder.On("Send", SuccessMail).Return(nil)
	mockBuilder.On("Send", ErrorMail).Return()

	return mockBuilder
}

func (s *MockSender) Send(mail entity.Mail) error {
	args := s.Called(mail)

	return args.Error(0)
}
