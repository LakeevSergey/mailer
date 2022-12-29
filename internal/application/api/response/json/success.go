package json

import (
	"encoding/json"

	"github.com/LakeevSergey/mailer/internal/application/api/response"
	"github.com/LakeevSergey/mailer/internal/application/api/response/text"
)

type successBody struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(data interface{}, status int) response.Response {
	body, err := json.Marshal(
		successBody{
			Status: status,
			Data:   data,
		},
	)
	if err != nil {
		return text.InternalErrorResponse()
	}
	return response.Response{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}
