package main

import (
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/server"
	"github.com/forecho/go-rest-api/internal/server/routes"
	"github.com/forecho/go-rest-api/pkg/logger"
	"os"
)

// @title Echo Demo App
// @version 1.0
// @description This is a demo version of Echo app.

// @contact.name forecho
// @contact.url https://forecho.com/
// @contact.email caizhenghai@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func main() {
	//// load application configurations
	cfg, err := config.Load()
	if err != nil {
		logger.Ins.Errorf("failed to load application configuration: %v", err)
		os.Exit(-1)
	}

	//if err = logger.Init(cfg); err != nil {
	//	log.Fatal().Msgf("Error initializing logger: '%v'", err)
	//	return
	//}

	// build HTTP server
	app := server.NewServer(cfg).CustomConfig()
	routes.ConfigureRoutes(app)
	app.Start()
}
