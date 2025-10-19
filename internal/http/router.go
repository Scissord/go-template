package http

import (
	"database/sql"
	"net/http"

	clientHandler "go-template/internal/http/client"
	clientRepo "go-template/internal/repository/client"
	clientService "go-template/internal/service/client"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// === client module ===
	repo := clientRepo.NewRepository(db)
	service := clientService.NewService(repo)
	handler := clientHandler.NewHandler(service)
	clientHandler.RegisterRoutes(r, handler)

	// healthcheck
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	return r
}
