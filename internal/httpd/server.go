package httpd

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/handlers"
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

	db := initDb(cfg)
	handlers.NewHandler(db)

	Start(cfg)
}

func initDb(cfg *config.Config) *ent.Client {
	drv, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.Error().Msgf("failed opening to mysql: '%v'", err)
		os.Exit(-1)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil {
		log.Error().Msgf("failed connection to mysql: '%v'", err)
		os.Exit(-1)
	}

	conn := ent.NewClient(ent.Driver(drv))

	defer func() {
		if err := conn.Close(); err != nil {
			log.Error().Err(err)
		}
	}()

	ctx := context.Background()
	if err := conn.Schema.Create(ctx); err != nil {
		log.Error().Err(err)
		return nil
	}
	log.Info().Msgf("DB Schema was created")
	return conn
}

// Start starts the echo HTTP server
func Start(cfg *config.Config) {
	e := echo.New()

	middleware.Register(e)
	handlers.Register(e)

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
