package sign

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPRequestSignChecker struct {
	hashChecker     HashChecker
	hashHeader      string
	includedHeaders []string
}

func NewHTTPRequestSignChecker(hashChecker HashChecker, hashHeader string, includedHeaders []string) *HTTPRequestSignChecker {
	return &HTTPRequestSignChecker{
		hashChecker:     hashChecker,
		hashHeader:      hashHeader,
		includedHeaders: includedHeaders,
	}
}

func (h *HTTPRequestSignChecker) Check(r *http.Request) bool {
	data, err := h.getData(r)
	if err != nil {
		return false
	}
	return h.hashChecker.Equal(data, r.Header.Get(h.hashHeader))
}

func (h *HTTPRequestSignChecker) getData(r *http.Request) (string, error) {
	var includedHeadersString string
	for _, key := range h.includedHeaders {
		includedHeadersString = includedHeadersString + fmt.Sprintf("\n%s:%s", key, r.Header.Get(key))
	}
	respBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	queueParams := r.URL.RawQuery
	if queueParams != "" {
		queueParams = "?" + queueParams
	}

	return fmt.Sprintf("%s:%s%s:%s%s", r.Method, r.URL.Path, queueParams, string(respBody), includedHeadersString), nil
}
