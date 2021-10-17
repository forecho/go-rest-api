package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v2"
	"os"
)

// Register middleware with echo
func Register(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	e.Use(middleware.CORS())
	//if viper.GetBool("cors-enabled") {
	//	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//		AllowOrigins:     viper.GetStringSlice("cors-allow-origins"),
	//		AllowMethods:     viper.GetStringSlice("cors-allow-methods"),
	//		AllowHeaders:     viper.GetStringSlice("cors-allow-headers"),
	//		AllowCredentials: viper.GetBool("cors-allow-credentials"),
	//		ExposeHeaders:    viper.GetStringSlice("cors-expose-headers"),
	//		MaxAge:           viper.GetInt("cors-max-age"),
	//	}))
	//}

	//if !viper.GetBool("log-requests-disabled") {
	//e.Use(middleware.Logger())
	logger := lecho.New(
		os.Stdout,
		lecho.WithTimestamp(),
		lecho.WithLevel(log.INFO),
		lecho.WithCallerWithSkipFrameCount(zerolog.CallerSkipFrameCount+1),
	)
	e.Logger = logger
	e.Use(lecho.Middleware(lecho.Config{
		Logger: logger,
	}))

}
