package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Welcome returns the welcome message
func (h *Handler) Welcome(c echo.Context) error {
	m := fmt.Sprintf("Welcome to echo,")

	return c.JSON(http.StatusOK, map[string]string{"message": m})
}
