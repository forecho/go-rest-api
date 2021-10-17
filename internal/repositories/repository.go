package repositories

import "github.com/forecho/go-rest-api/ent"

type Repository struct {
	DB *ent.Client
}

func NewRepository(db *ent.Client) *Repository {
	return &Repository{DB: db}
}
