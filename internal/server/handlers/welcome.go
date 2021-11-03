package handlers

import (
	"fmt"
	"github.com/forecho/go-rest-api/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Welcome returns the welcome message
func (h *Handler) Welcome(c echo.Context) error {
	m := fmt.Sprintf("Welcome to echo,")
	logger.With(c).Info("test111")
	return c.JSON(http.StatusOK, map[string]string{"message": m})
}
