package server

import (
	"net/http"

	"cache-server/internal/cache"
)

type Server struct {
	mux   *http.ServeMux
	cache *cache.Client
	// TODO: add logger
}

func New(cache *cache.Client) *Server {
	return &Server{
		mux:   http.NewServeMux(),
		cache: cache,
	}
}

// TODO: routes

func (s *Server) Run(port string) error {
	return http.ListenAndServe(port, s.mux)
}
