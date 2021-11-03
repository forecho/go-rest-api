package logger

import (
	"github.com/forecho/go-rest-api/internal/config"
	"github.com/forecho/go-rest-api/internal/constant"
	"github.com/forecho/go-rest-api/pkg/path"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

var Ins = logrus.New()

func init() {

	var level logrus.Level

	switch config.GetString("LOG_LEVEL") {
	case "warning":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
	case "fatal":
		level = logrus.FatalLevel
	case "panic":
		level = logrus.PanicLevel
	case "debug":
		level = logrus.DebugLevel
	default:
		level = logrus.InfoLevel
	}

	//Ins.SetFormatter(&LogFormatter{logrus.JSONFormatter{}})
	Ins.SetFormatter(&logrus.JSONFormatter{})
	Ins.SetOutput(os.Stdout)
	Ins.SetLevel(level)
	Ins.SetReportCaller(true)
	//local()
	Ins.AddHook(NewLfsHook(path.StoragePath()+"/logs/", "app"))
}

func With(c echo.Context) *logrus.Entry {
	fields := make(map[string]interface{})
	if c != nil {
		req := c.Request()
		res := c.Response()
		requestId := req.Header.Get(echo.HeaderXRequestID)
		if requestId == "" {
			requestId = res.Header().Get(echo.HeaderXRequestID)
		}
		fields[constant.RequestTraceID] = requestId
	}
	return Ins.WithFields(fields)
}
