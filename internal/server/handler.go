package server

import (
	"log"
	"net/http"
)

const key string = "url"

type Payload struct{}

type Response struct{}

func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get(key)

	if url == "" {
		log.Printf("a query parameter is required\n")
		s.BadRequest(w, "a query parameter is required")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get(key)

	if url == "" {
		log.Printf("a query parameter is required\n")
		s.BadRequest(w, "a query parameter is required")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
