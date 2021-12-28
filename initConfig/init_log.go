package initConfig

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"mikou/global"
	"mikou/pkg/logger"
)

func initLog() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
