package router

import "net/http"

type Api interface {
	Send() http.HandlerFunc
	AddTemplate() http.HandlerFunc
	GetTemplate() http.HandlerFunc
	SearchTemplates() http.HandlerFunc
	UpdateTemplate() http.HandlerFunc
	DeleteTemplate() http.HandlerFunc
}
