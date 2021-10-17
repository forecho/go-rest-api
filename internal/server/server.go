package server

import (
	"context"
	"fmt"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/db"
	_ "github.com/forecho/go-rest-api/pkg/auto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Echo   *echo.Echo
	DB     *ent.Client
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(cfg),
		Config: cfg,
	}
}

func (s *Server) CustomConfig() *Server {
	s.Echo.Validator = &CustomValidator{V: validator.New()}
	customErr := &customErrHandler{e: s.Echo}
	s.Echo.HTTPErrorHandler = customErr.handler
	s.Echo.Binder = &CustomBinder{b: &echo.DefaultBinder{}}
	return s
}

// Start server
func (s *Server) Start() {
	go func() {
		addr := fmt.Sprintf(":%d", s.Config.ServerPort)
		if err := s.Echo.Start(addr); err != nil {
			s.Echo.Logger.Info("Received signal, shutting down the server")
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	timeout := time.Duration(s.Config.GracefulTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
	}
	//return server.Echo.Start(":" + addr)
}

//
//func Init(cfg *mysql.Config) {
//	err := logging.Init(cfg)
//	if err != nil {
//		log.Fatal().Msgf("Error initializing logger: '%v'", err)
//	}
//
//	db := initDb(cfg)
//	handlers.NewHandler(db)
//
//	Start(cfg)
//}
//
//func initDb(cfg *mysql.Config) *ent.Client {
//	drv, err := sql.Open("mysql", cfg.DSN)
//	if err != nil {
//		log.Error().Msgf("failed opening to mysql: '%v'", err)
//		os.Exit(-1)
//	}
//	// Get the underlying sql.DB object of the driver.
//	db := drv.DB()
//	db.SetMaxIdleConns(10)
//	db.SetMaxOpenConns(100)
//	db.SetConnMaxLifetime(time.Hour)
//
//	err = db.Ping()
//	if err != nil {
//		log.Error().Msgf("failed connection to mysql: '%v'", err)
//		os.Exit(-1)
//	}
//
//	conn := ent.NewClient(ent.Driver(drv))
//
//	defer func() {
//		if err := conn.Close(); err != nil {
//			log.Error().Err(err)
//		}
//	}()
//
//	ctx := context.Background()
//	if err := conn.Schema.Create(ctx); err != nil {
//		log.Error().Err(err)
//		return nil
//	}
//	log.Info().Msgf("DB Schema was created")
//	return conn
//}
//
//// Start starts the echo HTTP server
//func Start(cfg *mysql.Config) {
//	e := echo.CustomConfig()
//
//	middleware.Register(e)
//	handlers.Register(e)
//
//	// Start server
//	go func() {
//		addr := fmt.Sprintf(":%d", cfg.ServerPort)
//		if err := e.Start(addr); err != nil {
//			e.Logger.Info("Received signal, shutting down the server")
//		}
//	}()
//
//	sig := make(chan os.Signal)
//	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
//	<-sig
//
//	timeout := time.Duration(cfg.GracefulTimeout) * time.Second
//	ctx, cancel := context.WithTimeout(context.Background(), timeout)
//	defer cancel()
//
//	if err := e.Shutdown(ctx); err != nil {
//		e.Logger.Fatal(err)
//	}
//}
