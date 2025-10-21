package product

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/products", h.GetAll).Methods("GET")
	r.HandleFunc("/products", h.Create).Methods("POST")
}
