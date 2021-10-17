package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/config"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ServiceWrapper interface {
	CreateAccessToken(user *ent.User) (accessToken string, exp int64, err error)
}

type Service struct {
	config *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
	}
}
