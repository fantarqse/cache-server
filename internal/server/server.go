package server

import (
	"fmt"
	"net/http"

	"cache-server/internal/cache"
)

type Server struct {
	mux   *http.ServeMux
	cache *cache.Client
}

func New(cache *cache.Client) *Server {
	return &Server{
		mux:   http.NewServeMux(),
		cache: cache,
	}
}

func (s *Server) Routes() {
	s.mux.HandleFunc("GET /items/top", logMiddleware(s.GetTop))
	s.mux.HandleFunc("GET /items", logMiddleware(s.Get))
	s.mux.HandleFunc("PUT /items", logMiddleware(s.Put))
	s.mux.HandleFunc("DELETE /items", logMiddleware(s.Delete))
}

func (s *Server) Run(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.mux)
}
