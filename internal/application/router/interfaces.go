package router

import "net/http"

type Api interface {
	Send() http.HandlerFunc
}
