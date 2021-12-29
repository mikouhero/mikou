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
		MaxSize:   global.AppSetting.LogMaxSize,
		MaxAge:    global.AppSetting.LogMaxAge,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}


