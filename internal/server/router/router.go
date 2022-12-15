package router

import (
	"log"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/server/api/response/text"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(api Api, logger *log.Logger) chi.Router {
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger, NoColor: true})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.HandleFunc("/*", error404)
	r.Post("/send", api.Send())

	return r
}

func error404(rw http.ResponseWriter, r *http.Request) {
	text.Error404Response().Write(rw)
}
