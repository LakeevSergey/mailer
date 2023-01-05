package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/json"
)

type TimestampChecker struct {
	header string
	delta  time.Duration
}

func NewTimestampChecker(header string, delta time.Duration) *TimestampChecker {
	return &TimestampChecker{
		header: header,
		delta:  delta,
	}
}

func (c *TimestampChecker) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.checkTimestamp(r) {
			next.ServeHTTP(w, r)
		} else {
			json.ErrorResponse(http.StatusText(http.StatusForbidden), http.StatusForbidden).Write(w)
		}
	})
}

func (c *TimestampChecker) checkTimestamp(r *http.Request) bool {
	current := time.Now()
	fromHeader, err := strconv.ParseInt(r.Header.Get(c.header), 10, 64)
	if err != nil {
		return false
	}
	return current.Sub(time.Unix(fromHeader, 0)).Abs() < c.delta
}
