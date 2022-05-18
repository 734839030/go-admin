package middleware

import (
	"github.com/sirupsen/logrus"
	"go-admin/configs"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {
	level, err := logrus.ParseLevel(configs.C.Log.Level)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)

	logger := &lumberjack.Logger{
		Filename:   configs.C.Log.Filename,
		MaxSize:    configs.C.Log.MaxSize,
		MaxAge:     configs.C.Log.MaxAge,
		MaxBackups: 0,
		LocalTime:  true,
	}
	logrus.SetReportCaller(configs.C.Log.ReportCaller)
	if !configs.C.Log.OutputToConsole {
		logrus.SetOutput(logger)
	}

}
