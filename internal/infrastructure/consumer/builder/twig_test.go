package builder

import (
	"testing"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	storagermock "github.com/LakeevSergey/mailer/internal/infrastructure/storager/mock"
	"github.com/stretchr/testify/assert"
)

func TestTwigBuilder_Build(t *testing.T) {
	type args struct {
		template entity.Template
		params   map[string]string
	}
	tests := []struct {
		name      string
		b         *TwigBuilder
		args      args
		wantBody  string
		wantTitle string
		wantErr   error
	}{
		{
			name: "simple test",
			b:    &TwigBuilder{},
			args: args{
				template: storagermock.ActiveTemplate,
				params:   map[string]string{"name": "Name"},
			},
			wantBody:  "Active template body, Name",
			wantTitle: "Active template title, Name",
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, gotTitle, err := tt.b.Build(tt.args.template, tt.args.params)

			assert.ErrorIs(t, tt.wantErr, err)
			assert.Equal(t, tt.wantBody, gotBody)
			assert.Equal(t, tt.wantTitle, gotTitle)
		})
	}
}
