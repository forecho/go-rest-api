package db

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/forecho/go-rest-api/ent"
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/pkg/logger"
	"os"
	"time"
)

func Init(cfg *config.Config) *ent.Client {

	drv, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		logger.Ins.Errorf("failed opening to mysql: '%v'", err)
		os.Exit(-1)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	err = db.Ping()
	if err != nil {
		logger.Ins.Errorf("failed connection to mysql: '%v'", err)
		os.Exit(-1)
	}

	conn := ent.NewClient(ent.Driver(drv))

	defer func() {
		if err := conn.Close(); err != nil {
			logger.Ins.Errorf("failed close to mysql: '%v'", err)
		}
	}()

	ctx := context.Background()
	if err := conn.Schema.Create(ctx); err != nil {
		logger.Ins.Errorf("failed create schema to mysql: '%v'", err)
		return nil
	}
	logger.Ins.Info("DB Schema was created")
	return conn
}
