package requestprocessor

import (
	"context"
	"testing"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/infrastructure/consumer/builder"
	"github.com/LakeevSergey/mailer/internal/infrastructure/consumer/sender"
	"github.com/LakeevSergey/mailer/internal/infrastructure/storager/mock"
	"github.com/stretchr/testify/assert"
)

func TestSendMailRequestProcessor_Process(t *testing.T) {
	ctx := context.Background()
	mockStorager := mock.NewMockStorager(ctx)
	mockBuilder := builder.NewMockBuilder()
	mockSender := sender.NewMockSender()

	type args struct {
		ctx      context.Context
		sendMail entity.SendMail
	}
	tests := []struct {
		name string
		p    *SendMailRequestProcessor
		args args
		err  error
	}{
		{
			name: "simple test",
			p: &SendMailRequestProcessor{
				storager: mockStorager,
				builder:  mockBuilder,
				sender:   mockSender,
				defaultSendFrom: entity.SendFrom{
					Name:  "test",
					Email: "test@test.test",
				},
			},
			args: args{
				ctx: context.Background(),
				sendMail: entity.SendMail{
					Code:   "active_template",
					SendTo: []string{"test@test.test"},
					Params: map[string]string{"name": "Name"},
				},
			},
			err: nil,
		},
		{
			name: "simple test",
			p: &SendMailRequestProcessor{
				storager: mockStorager,
				builder:  mockBuilder,
				sender:   mockSender,
				defaultSendFrom: entity.SendFrom{
					Name:  "test",
					Email: "test@test.test",
				},
			},
			args: args{
				ctx: context.Background(),
				sendMail: entity.SendMail{
					Code:   "not_active_template",
					SendTo: []string{"test@test.test"},
					Params: map[string]string{"name": "Name"},
				},
			},
			err: ErrorTemplateDeactivated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.Process(tt.args.ctx, tt.args.sendMail)
			if tt.err != nil {
				assert.ErrorIs(t, tt.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
