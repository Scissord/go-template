package client

import (
	repo "go-template/internal/repository/product"
)

type Service struct {
	repo *repo.Repository
}

func NewService(r *repo.Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAllClients() ([]repo.Client, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateClient(name, email string) (*repo.Client, error) {
	client := &repo.Client{Name: name, Email: email}
	err := s.repo.Create(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}
