package global

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"mikou/pkg/logger"
	"mikou/pkg/setting"
)

var (
	ServerSetting     *setting.ServerSettingS
	AppSetting        *setting.AppSettingS
	DatabaseSetting   *setting.DatabaseSettingS
	DatabaseSettingV2 *setting.DatabaseSettingSV2
	Logger            *logger.Logger
	LoggerV2          *zap.SugaredLogger
	TokenSetting      *setting.TokenSettingS
	Cron              *cron.Cron
)
