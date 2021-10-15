package httpd

import (
	"context"
	"fmt"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/config"
	handlers2 "github.com/forecho/go-rest-api/internal/handlers"
	"github.com/forecho/go-rest-api/internal/httpd/middleware"
	_ "github.com/forecho/go-rest-api/pkg/auto"
	"github.com/forecho/go-rest-api/pkg/logging"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init(cfg *config.Config) {
	err := logging.Init(cfg)
	if err != nil {
		log.Fatal().Msgf("Error initializing logger: '%v'", err)
	}

	// connect to the database
	db, err := ent.Open("mysql", cfg.DSN)
	if err != nil {
		log.Error().Err(err)
		os.Exit(-1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err)
		}
	}()

	Start(cfg)
}

// Start starts the echo HTTP server
func Start(cfg *config.Config) {
	e := echo.New()

	middleware.Register(e)
	handlers2.Register(e)

	// Start server
	go func() {
		addr := fmt.Sprintf(":%d", cfg.ServerPort)
		if err := e.Start(addr); err != nil {
			e.Logger.Info("Received signal, shutting down the server")
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	timeout := time.Duration(cfg.GracefulTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
