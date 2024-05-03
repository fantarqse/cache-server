package server

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (s *Server) Error(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(errorResponse{
		Error: msg,
		Code:  code,
	})
}

func (s *Server) BadRequest(w http.ResponseWriter, msg string) {
	s.Error(w, msg, http.StatusBadRequest)
}

func (s *Server) NotFound(w http.ResponseWriter, msg string) {
	s.Error(w, msg, http.StatusNotFound)
}

func (s *Server) InternalServerError(w http.ResponseWriter, msg string) {
	s.Error(w, msg, http.StatusInternalServerError)
}
