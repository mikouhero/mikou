package initConfig

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func initlog2() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	LoggerV2 := zap.New(core, zap.AddCaller(),zap.AddStacktrace(zap.ErrorLevel))
	global.LoggerV2 = LoggerV2.Sugar()
}



func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   global.AppSetting.LogMaxSize,
		MaxAge:    global.AppSetting.LogMaxAge,
		LocalTime: true,
		Compress:  false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
