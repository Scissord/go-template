package http

import (
	"database/sql"
	"log"
	"net/http"

	client "go-template/internal/http/client"
	product "go-template/internal/http/product"

	"github.com/gorilla/mux"
)

type Module interface {
	RegisterRoutes(r *mux.Router)
}

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	modules := []Module{
		client.NewModule(db),
		product.NewModule(db),
	}

	for _, m := range modules {
		m.RegisterRoutes(r)
	}

	// Health-check
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Printf("error writing response: %v", err)
		}
	}).Methods("GET")

	return r
}
