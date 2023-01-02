package json

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/domain/mailsender"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager"
)

type JSONApi struct {
	sender          mailsender.MailSender
	templateManager templatemanager.TemplateManager
}

func NewJSONApi(sender mailsender.MailSender, templateManager templatemanager.TemplateManager) *JSONApi {
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
