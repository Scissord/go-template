package client

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]Client, error) {
	rows, err := r.DB.Query("SELECT id, name, email, created_at FROM client ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func (r *Repository) Create(c *Client) error {
	return r.DB.QueryRow(
		"INSERT INTO client (name, email) VALUES ($1, $2) RETURNING id, created_at",
		c.Name, c.Email,
	).Scan(&c.ID, &c.CreatedAt)
}
