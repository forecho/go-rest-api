package middleware

import (
	"github.com/forecho/go-rest-api/pkg/logger"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

// Logger returns a middleware that logs HTTP requests.
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			p := req.URL.Path
			if p == "" {
				p = "/"
			}

			bytesIn := req.Header.Get(echo.HeaderContentLength)
			if bytesIn == "" {
				bytesIn = "0"
			}

			logger.With(c).WithFields(map[string]interface{}{
				"time_rfc3339":  time.Now().Format(time.RFC3339),
				"remote_ip":     c.RealIP(),
				"host":          req.Host,
				"uri":           req.RequestURI,
				"method":        req.Method,
				"path":          p,
				"referer":       req.Referer(),
				"user_agent":    req.UserAgent(),
				"status":        res.Status,
				"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
				"latency_human": stop.Sub(start).String(),
				"bytes_in":      bytesIn,
				"bytes_out":     strconv.FormatInt(res.Size, 10),
			}).Info("Handled request")

			return nil
		}
	}
}
