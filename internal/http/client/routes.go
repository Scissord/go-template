package client

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, h *Handler) {
	s := r.PathPrefix("/clients").Subrouter()
	s.HandleFunc("", h.GetAll).Methods("GET")
	s.HandleFunc("", h.Create).Methods("POST")
}
