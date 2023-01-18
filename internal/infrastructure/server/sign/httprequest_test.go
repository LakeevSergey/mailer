package sign

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/LakeevSergey/mailer/internal/infrastructure/server/hasher"
)

func TestHTTPRequestSignChecker_Check(t *testing.T) {
	mockHasher := hasher.NewMockHasher()

	type args struct {
		method  string
		url     string
		body    string
		headers map[string]string
	}
	tests := []struct {
		name string
		h    *HTTPRequestSignChecker
		args args
		want bool
	}{
		{
			name: "simple test",
			h: &HTTPRequestSignChecker{
				hashChecker:     mockHasher,
				hashHeader:      "Signature",
				includedHeaders: []string{"Timestamp"},
			},
			args: args{
				method: "POST",
				url:    "http://host/path?param=value",
				body:   "body",
				headers: map[string]string{
					"Timestamp": "1234",
					"Signature": "POST:/path?param=value:body\nTimestamp:1234",
				},
			},
			want: true,
		},
		{
			name: "negative test",
			h: &HTTPRequestSignChecker{
				hashChecker:     mockHasher,
				hashHeader:      "Signature",
				includedHeaders: []string{"Timestamp"},
			},
			args: args{
				method: "POST",
				url:    "http://host/path?param=value",
				body:   "body",
				headers: map[string]string{
					"Timestamp": "1234",
					"Signature": "POST:/path:body\nTimestamp:1234",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := new(bytes.Buffer)
			_, err := body.Write([]byte(tt.args.body))
			assert.NoError(t, err)

			request, err := http.NewRequest(tt.args.method, tt.args.url, body)
			assert.NoError(t, err)

			for header, val := range tt.args.headers {
				request.Header.Add(header, val)
			}

			assert.Equal(t, tt.want, tt.h.Check(request))
		})
	}
}
