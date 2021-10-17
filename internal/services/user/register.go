package user

import (
	"context"
	"github.com/forecho/go-rest-api/internal/requests"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(r *requests.RegisterRequest) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(r.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return
	}
	ctx := context.Background()
	_, err = s.DB.User.
		Create().
		SetPassword(string(encryptedPassword)).
		SetUsername(r.Username).
		SetEmail(r.Email).
		Save(ctx)
	return
}
