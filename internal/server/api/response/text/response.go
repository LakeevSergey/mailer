package text

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/server/api/response"
)

func Response(text string, status int) response.Response {
	return response.Response{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       []byte(text),
	}
}

func Error404Response() response.Response {
	return Response(http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func InternalErrorResponse() response.Response {
	return Response(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
