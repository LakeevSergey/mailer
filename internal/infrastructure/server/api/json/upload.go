package json

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/domain/attachmentmanager/dto"
	apijsondto "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/dto"
	responsejson "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
)

func (a *JSONApi) Upload() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("attachment")
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}
		defer file.Close()

		dtoAdd := dto.Add{
			FileName: handler.Filename,
			Mime:     handler.Header.Get("Content-Type"),
			Content:  file,
		}

		id, size, err := a.attachmentManager.Add(r.Context(), dtoAdd)
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}

		responsejson.SuccessResponse(apijsondto.FileInfo{
			Id:   id,
			Size: size,
		}, http.StatusCreated).Write(rw)
	}
}
