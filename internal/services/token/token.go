package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/forecho/go-rest-api/ent"
	"time"
)

func (s *Service) CreateAccessToken(user *ent.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Duration(s.config.JWTExpiration) * time.Hour).Unix()
	claims := &JwtCustomClaims{
		user.Username,
		uint(user.ID),
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.config.JWTSigningKey))
	if err != nil {
		return "", 0, err
	}

	return t, exp, err
}
