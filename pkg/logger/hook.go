package logger

import (
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

// NewLfsHook WithMaxAge和WithRotationCount二者只能设置一个
func NewLfsHook(filePath, logName string) logrus.Hook {
	infoWriter, err := rotate.New(
		// 分割后的文件名称
		filePath+logName+".%Y%m%d%H.log",
		// 生成软链，指向最新日志文件
		rotate.WithLinkName(filePath+logName),
		// 设置日志切割时间间隔(1天)
		rotate.WithRotationTime(24*time.Hour),
		// 设置最大保存时间(30天)
		// rotate.WithMaxAge(30*24*time.Hour),
		// 设置文件清理前最多保存的个数
		rotate.WithRotationCount(50),
	)
	errorWriter, err := rotate.New(
		// 分割后的文件名称
		filePath+logName+".%Y%m%d%H.error.log",
		// 生成软链，指向最新日志文件
		rotate.WithLinkName(filePath+logName),
		// 设置日志切割时间间隔(1天)
		rotate.WithRotationTime(24*time.Hour),
		// 设置最大保存时间(30天)
		// rotate.WithMaxAge(30*24*time.Hour),
		// 设置文件清理前最多保存的个数
		rotate.WithRotationCount(50),
	)
	warnWriter, err := rotate.New(
		// 分割后的文件名称
		filePath+logName+".%Y%m%d%H.warn.log",
		// 生成软链，指向最新日志文件
		rotate.WithLinkName(filePath+logName),
		// 设置日志切割时间间隔(1天)
		rotate.WithRotationTime(24*time.Hour),
		// 设置最大保存时间(30天)
		// rotate.WithMaxAge(30*24*time.Hour),
		// 设置文件清理前最多保存的个数
		rotate.WithRotationCount(50),
	)

	if err != nil {
		logrus.Errorf("config logger error: %v", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: infoWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: errorWriter,
		logrus.PanicLevel: errorWriter,
	}

	lfsHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})

	return lfsHook
}
