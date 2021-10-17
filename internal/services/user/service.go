package user

import (
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/requests"
)

type ServiceWrapper interface {
	Register(r *requests.RegisterRequest) (err error)
}

type Service struct {
	DB *ent.Client
}

func NewService(db *ent.Client) *Service {
	return &Service{
		DB: db,
	}
}
