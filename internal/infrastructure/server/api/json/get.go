package json

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LakeevSergey/mailer/internal/domain"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/dto"
	responsejson "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
	"github.com/go-chi/chi/v5"
)

func (a *JSONApi) GetTemplate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		}
		template, err := a.templateManager.Get(r.Context(), id)

		if errors.Is(err, domain.ErrorTemplateNotFound) {
			responsejson.ErrorResponse(err.Error(), http.StatusNotFound).Write(rw)
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
