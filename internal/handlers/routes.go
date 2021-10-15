package handlers

import (
	"github.com/forecho/go-rest-api/internal/httpd/router"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

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
		//{"Healthz", http.MethodGet, "/healthz", handler.Healthz}},
	},
}
