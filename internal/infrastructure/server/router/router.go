package router

import (
	"net/http"

	"github.com/LakeevSergey/mailer/internal/infrastructure"
	"github.com/LakeevSergey/mailer/internal/infrastructure/server/api/response/text"
	appmiddleware "github.com/LakeevSergey/mailer/internal/infrastructure/server/router/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(api Api, logger infrastructure.Logger, middlewares ...func(next http.Handler) http.Handler) chi.Router {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: NewLoggerAdapter(logger), NoColor: true})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(appmiddleware.Compressor)
	r.Use(appmiddleware.Decompressor)

	for _, m := range middlewares {
		r.Use(m)
	}

	r.HandleFunc("/*", error404)
	r.Get("/template", api.SearchTemplates())
	r.Post("/template", api.AddTemplate())
	r.Get("/template/{id}", api.GetTemplate())
	r.Post("/template/{id}", api.UpdateTemplate())
	r.Delete("/template/{id}", api.DeleteTemplate())
	r.Post("/send", api.Send())
	r.Post("/upload", api.Upload())

	return r
}

func error404(rw http.ResponseWriter, r *http.Request) {
	text.Error404Response().Write(rw)
}
