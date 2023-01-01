package json

import (
	"errors"
	"net/http"
	"strconv"

	responsejson "github.com/LakeevSergey/mailer/internal/application/api/response/json"
	"github.com/LakeevSergey/mailer/internal/domain"
	"github.com/go-chi/chi/v5"
)

func (a *JSONApi) DeleteTemplate() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		}

		err = a.templateManager.Delete(r.Context(), id)
		if errors.Is(err, domain.ErrorTemplateNotFound) {
			responsejson.ErrorResponse(err.Error(), http.StatusNotFound).Write(rw)
			return
		} else if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}

		responsejson.SuccessResponse("Ok", http.StatusOK).Write(rw)
	}
}
