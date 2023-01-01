package router

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/application"
	"github.com/LakeevSergey/mailer/internal/application/api/response/text"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(api Api, logger application.Logger) chi.Router {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: NewLoggerAdaptor(logger), NoColor: true})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.HandleFunc("/*", error404)
	r.Get("/template", api.SearchTemplates())
	r.Post("/template", api.AddTemplate())
	r.Get("/template/{id}", api.GetTemplate())
	r.Post("/template/{id}", api.UpdateTemplate())
	r.Delete("/template/{id}", api.DeleteTemplate())
	r.Post("/send", api.Send())

	return r
}

func error404(rw http.ResponseWriter, r *http.Request) {
	text.Error404Response().Write(rw)
}
