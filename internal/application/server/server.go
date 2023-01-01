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

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}
	errCh := make(chan error)

	go func() {
		s.logger.Info("Server is running")
		errCh <- srv.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		srv.Shutdown(ctx)
		return nil
	}
}
