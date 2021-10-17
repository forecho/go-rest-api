package routes

import (
	_ "github.com/forecho/go-rest-api/docs"
	"github.com/forecho/go-rest-api/internal/server"
	"github.com/forecho/go-rest-api/internal/server/handlers"
	"github.com/forecho/go-rest-api/internal/server/middleware"
	"github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(s *server.Server) {
	middleware.Register(s.Echo)

	handler := handlers.RegisterHandlers(s)

	s.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	s.Echo.POST("/login", handler.Login)
	s.Echo.POST("/register", handler.Register)

	r := s.Echo.Group("")
	//config := middleware.JWTConfig{
	//	Claims:     &token.JwtCustomClaims{},
	//	SigningKey: []byte(s.Config.Auth.AccessSecret),
	//}
	//r.Use(middleware.JWTWithConfig(config))

	r.GET("/", handler.Welcome)
}
