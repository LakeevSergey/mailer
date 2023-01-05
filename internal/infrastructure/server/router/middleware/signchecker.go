package middleware

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
)

type SignCheckerMiddleware struct {
	signChecker SignChecker
}

func NewSignCheckerMiddleware(signChecker SignChecker) *SignCheckerMiddleware {
	return &SignCheckerMiddleware{
		signChecker: signChecker,
	}
}

func (m *SignCheckerMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.signChecker.Check(r) {
			next.ServeHTTP(w, r)
			return
		} else {
			json.ErrorResponse(http.StatusText(http.StatusForbidden), http.StatusForbidden).Write(w)
			return
		}
	})
}
