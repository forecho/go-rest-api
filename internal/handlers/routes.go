package handlers

import (
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/httpd/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	db *ent.Client
}

func NewHandler(client *ent.Client) *Handler {
	return &Handler{db: client}
}

var handler Handler

// Register routes with echo
func Register(e *echo.Echo) *Handler {
	for _, route := range Router.Routes {
		e.Add(route.Method, route.Pattern, route.HandlerFunc)
	}

	return &handler
}

var Router = &router.Router{
	Routes: []router.Route{
		{"Welcome", http.MethodGet, "/", handler.Welcome},
		{"Create User", http.MethodPost, "/users", handler.CreateUser},
	},
}
