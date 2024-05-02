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

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /items/top", s.GetTop)
	s.mux.HandleFunc("GET /items", s.Get)
	s.mux.HandleFunc("PUT /items", s.Put)
	s.mux.HandleFunc("DELETE /items", s.Delete)
}

func (s *Server) Run(port string) error {
	return http.ListenAndServe(port, s.mux)
}
