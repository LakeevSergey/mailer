package json

import (
	"encoding/json"

	"github.com/LakeevSergey/mailer/internal/server/api/response"
	"github.com/LakeevSergey/mailer/internal/server/api/response/text"
)

type errorBody struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func ErrorResponse(error string, status int) response.Response {
	body, err := json.Marshal(errorBody{
		Status: status,
		Error:  error,
	})
	if err != nil {
		return text.InternalErrorResponse()
	}
	return response.Response{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}
