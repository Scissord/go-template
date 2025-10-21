package product

import (
	"database/sql"

	productRepo "go-template/internal/repository/product"
	productService "go-template/internal/service/product"

	"github.com/gorilla/mux"
)

type Module struct {
	db *sql.DB
}

func NewModule(db *sql.DB) *Module {
	return &Module{db: db}
}

func (m *Module) RegisterRoutes(r *mux.Router) {
	repo := productRepo.NewRepository(m.db)
	service := productService.NewService(repo)
	handler := NewHandler(service)
	RegisterRoutes(r, handler)
}
