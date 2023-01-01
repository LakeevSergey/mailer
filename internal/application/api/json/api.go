package json

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/application/api"
)

type JSONApi struct {
	sender          api.MailSender
	templateManager api.TemplateManager
}

func NewJSONApi(sender api.MailSender, templateManager api.TemplateManager) *JSONApi {
	return &JSONApi{
		sender:          sender,
		templateManager: templateManager,
	}
}

func unserializeRequest[T any](r *http.Request) (T, error) {
	var result T
	respBody, err := io.ReadAll(r.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
