package main

import (
	"flag"
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/httpd"
	"github.com/rs/zerolog/log"
	"os"
)

var flagConfig = flag.String("config", "./config/local.yml", "path to the config file")

func main() {
	flag.Parse()
	//// load application configurations
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Error().Msgf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

	// build HTTP server
	httpd.Init(cfg)
}
