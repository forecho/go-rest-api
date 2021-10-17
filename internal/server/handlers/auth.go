package handlers

import (
	"github.com/forecho/go-rest-api/internal/repositories"
	"github.com/forecho/go-rest-api/internal/requests"
	"github.com/forecho/go-rest-api/internal/responses"
	"github.com/forecho/go-rest-api/internal/services/token"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Login godoc
// @Summary Authenticate a user
// @Description Perform user login
// @ID user-login
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.LoginRequest true "User's credentials"
// @Success 200 {object} responses.LoginResponse
// @Failure 401 {object} responses.Error
// @Router /login [post]
func (h *Handler) Login(c echo.Context) error {
	request := new(requests.LoginRequest)

	if err := c.Bind(request); err != nil {
		return err
	}

	u := repositories.NewRepository(h.server.DB).GetUserByEmail(request.Email)
	h.server.Echo.Logger.Infof("1111111111: %v", u)
	if u == nil || (bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(request.Password)) != nil) {
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	}

	accessToken, exp, err := token.NewService(h.server.Config).CreateAccessToken(u)
	if err != nil {
		return err
	}
	res := responses.NewLoginResponse(accessToken, exp)

	return responses.Response(c, http.StatusOK, res)
}
