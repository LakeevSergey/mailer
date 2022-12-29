package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	address string
	handler http.Handler
}

func NewServer(address string, handler http.Handler) *Server {
	return &Server{
		address: address,
		handler: handler,
	}
}

func (s *Server) Run(ctx context.Context) {
	srv := &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}
	go log.Fatal(srv.ListenAndServe())
	<-ctx.Done()
	srv.Shutdown(ctx)
}
