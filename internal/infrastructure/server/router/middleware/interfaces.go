package middleware

import "net/http"

type SignChecker interface {
	Check(r *http.Request) bool
}
