package json

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/domain/entity"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/dto"
	responsejson "github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
)

func (a *JSONApi) Send() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		dto, err := unserializeRequest[dto.SendMail](r)
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusBadRequest).Write(rw)
			return
		}

		sendMail := entity.SendMail{
			Code:   dto.Code,
			SendTo: dto.SendTo,
			Params: dto.Params,
			Title:  dto.Title,
			Body:   dto.Body,
		}
		if dto.SendFrom != nil {
			sendMail.SendFrom = &entity.SendFrom{
				Name:  dto.SendFrom.Name,
				Email: dto.SendFrom.Email,
			}
		}

		err = a.sender.Send(sendMail)
		if err != nil {
			responsejson.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(rw)
			return
		}

		responsejson.SuccessResponse(http.StatusText(http.StatusOK), http.StatusOK).Write(rw)
	}
}
