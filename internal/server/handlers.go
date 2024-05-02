package server

import (
	"fmt"
	"net/http"
)

func (s *Server) GetTop(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		fmt.Sprintf("%s %s\n", r.Method, r.URL.String()),
	))
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		fmt.Sprintf("%s %s\n", r.Method, r.URL.String()),
	))
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		fmt.Sprintf("%s %s\n", r.Method, r.URL.String()),
	))
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		fmt.Sprintf("%s %s\n", r.Method, r.URL.String()),
	))
}
