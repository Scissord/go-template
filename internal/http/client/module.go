package client

import (
	"database/sql"

	clientRepo "go-template/internal/repository/client"
	clientService "go-template/internal/service/client"

	"github.com/gorilla/mux"
)

type Module struct {
	db *sql.DB
}

func NewModule(db *sql.DB) *Module {
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(r *mux.Router) {
	repo := clientRepo.NewRepository(m.db)
	service := clientService.NewService(repo)
	handler := NewHandler(service)

	RegisterRoutes(r, handler)
}
