package api

import (
	"io"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/server/api/response/json"
)

type JSONApi struct {
	decoder Decoder
	sender  MailSender
}

func NewJSONApi(decoder Decoder, sender MailSender) *JSONApi {
	return &JSONApi{}
}

func (a *JSONApi) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		respBody, err := io.ReadAll(r.Body)
		if err != nil {
			json.ErrorResponse(err.Error(), http.StatusBadRequest).Write(w)
			return
		}
		sendMail, err := a.decoder.Decode(respBody)
		if err != nil {
			json.ErrorResponse(err.Error(), http.StatusBadRequest).Write(w)
			return
		}

		err = a.sender.Send(sendMail)
		if err != nil {
			json.ErrorResponse(err.Error(), http.StatusInternalServerError).Write(w)
			return
		}

		json.SuccessResponse(http.StatusText(http.StatusOK), http.StatusOK).Write(w)
	}
}
