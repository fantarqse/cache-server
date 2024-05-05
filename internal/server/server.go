package server

import (
	"context"
	"fmt"
	"net/http"
)

type Storage interface {
	GetAll(ctx context.Context) error
	Get(ctx context.Context) error
	Put(ctx context.Context) error
	Delete(ctx context.Context) error
}

type Server struct {
	mux     *http.ServeMux
	storage Storage
}

func New(storage Storage) *Server {
	return &Server{
		mux:     http.NewServeMux(),
		storage: storage,
	}
}

func (s *Server) Run(port string) error {
	s.mux.HandleFunc("GET /items/top", logMiddleware(s.GetAll))
	s.mux.HandleFunc("GET /items", logMiddleware(s.Get))
	s.mux.HandleFunc("PUT /items", logMiddleware(s.Put))
	s.mux.HandleFunc("DELETE /items", logMiddleware(s.Delete))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.mux)
}
