package response

import (
	"net/http"
)

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func (r Response) Write(rw http.ResponseWriter) {
	for key, value := range r.Headers {
		rw.Header().Add(key, value)
	}
	rw.WriteHeader(r.StatusCode)
	rw.Write(r.Body)
}
