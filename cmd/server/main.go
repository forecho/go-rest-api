package main

import (
	"flag"
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/server"
	"github.com/forecho/go-rest-api/internal/server/routes"
	"github.com/rs/zerolog/log"
	"os"
)

var flagConfig = flag.String("mysql", "./config/local.yml", "path to the mysql file")

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
	flag.Parse()
	//// load application configurations
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Error().Msgf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// build HTTP server
	app := server.NewServer(cfg).CustomConfig()
	routes.ConfigureRoutes(app)
	app.Start()
}
