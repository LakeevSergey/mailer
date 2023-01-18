package templatemanager

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/LakeevSergey/mailer/internal/domain"
	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/infrastructure/storager/mock"
)

func TestTemplateManager_Get(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		m       *TemplateManager
		args    args
		want    entity.Template
		wantErr error
	}{
		{
			name: "simple test",
			m: &TemplateManager{
				storager: mock.NewMockStorager(ctx),
			},
			args: args{
				ctx,
				1,
			},
			want:    mock.ActiveTemplate,
			wantErr: nil,
		},
		{
			name: "template not found test",
			m: &TemplateManager{
				storager: mock.NewMockStorager(ctx),
			},
			args: args{
				ctx,
				3,
			},
			want:    entity.Template{},
			wantErr: domain.ErrorTemplateNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Get(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
