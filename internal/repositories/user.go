package repositories

import (
	"context"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/ent/user"
)

func (r *Repository) GetUserByEmail(email string) (u *ent.User) {
	ctx := context.Background()
	u, _ = r.DB.User.Query().Where(user.Email(email)).Only(ctx)
	return
}
