package middleware

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"trojan-panel-core/model/constant"
)

func InitLog() {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   constant.SystemLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetLevel(logrus.WarnLevel)
}
