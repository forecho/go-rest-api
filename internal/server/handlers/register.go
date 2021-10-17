package handlers

import (
	"github.com/forecho/go-rest-api/internal/repositories"
	"github.com/forecho/go-rest-api/internal/requests"
	"github.com/forecho/go-rest-api/internal/responses"
	"github.com/forecho/go-rest-api/internal/services/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Register godoc
// @Summary Register
// @Description New user registration
// @ID user-register
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.RegisterRequest true "User's email, user's password"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /register [post]
func (h *Handler) Register(c echo.Context) error {

	request := new(requests.RegisterRequest)

	if err := c.Bind(request); err != nil {
		return err
	}

	u := repositories.NewRepository(h.server.DB).GetUserByEmail(request.Email)
	if u.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	err := user.NewService(h.server.DB).Register(request)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	}

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}
