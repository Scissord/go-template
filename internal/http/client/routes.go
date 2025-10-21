package client

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/clients", h.GetAll).Methods("GET")
	r.HandleFunc("/clients", h.Create).Methods("POST")
}
