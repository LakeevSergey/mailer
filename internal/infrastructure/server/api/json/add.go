package json

import (
	"errors"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/domain"
	templatemanagerdto "github.com/LakeevSergey/mailer/internal/domain/templatemanager/dto"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/dto"
	responsejson "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
)

func (a *JSONApi) AddTemplate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		dtoTemplate, err := unserializeRequest[dto.Template](r)
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		}

		template, err := a.templateManager.Add(r.Context(), templatemanagerdto.Add{
			Active: dtoTemplate.Active,
			Code:   dtoTemplate.Code,
			Name:   dtoTemplate.Name,
			Body:   dtoTemplate.Body,
			Title:  dtoTemplate.Title,
		})
		if errors.Is(err, domain.ErrorTemplateCodeDuplicate) {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		} else if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}

		responsejson.SuccessResponse(dto.Template{
			Id:     template.Id,
			Active: template.Active,
			Code:   template.Code,
			Name:   template.Name,
			Body:   template.Body,
			Title:  template.Title,
		}, http.StatusCreated).Write(rw)
	}
}
