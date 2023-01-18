package json

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/domain/attachmentmanager"
	"github.com/LakeevSergey/mailer/internal/domain/mailsender"
	"github.com/LakeevSergey/mailer/internal/domain/templatemanager"
)

type JSONApi struct {
	sender            mailsender.MailSender
	templateManager   templatemanager.TemplateManager
	attachmentManager attachmentmanager.AttachmentManager
}

func NewJSONApi(sender mailsender.MailSender, templateManager templatemanager.TemplateManager, attachmentManager attachmentmanager.AttachmentManager) *JSONApi {
	return &JSONApi{
		sender:            sender,
		templateManager:   templateManager,
		attachmentManager: attachmentManager,
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
