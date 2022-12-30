package server

import (
	"context"
	"net/http"

	"github.com/LakeevSergey/mailer/internal/application"
)

type Server struct {
	address string
	handler http.Handler
	logger  application.Logger
}

func NewServer(address string, handler http.Handler, logger application.Logger) *Server {
	return &Server{
		address: address,
		handler: handler,
		logger:  logger,
	}
}

func (s *Server) Run(ctx context.Context) {
	srv := &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}
	go func() {
		s.logger.ErrorErr(srv.ListenAndServe())
	}()
	s.logger.Info("Server running")
	<-ctx.Done()
	srv.Shutdown(ctx)
}
