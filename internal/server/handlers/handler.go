package handlers

import (
	"github.com/forecho/go-rest-api/internal/server"
)

type Handler struct {
	server *server.Server
}

func RegisterHandlers(s *server.Server) *Handler {
	return &Handler{server: s}
}
