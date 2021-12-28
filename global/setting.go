package global

import (
	"mikou/pkg/logger"
	"mikou/pkg/setting"
)

var (
	ServerSetting     *setting.ServerSettingS
	AppSetting        *setting.AppSettingS
	DatabaseSetting   *setting.DatabaseSettingS
	DatabaseSettingV2 *setting.DatabaseSettingSV2
	Logger            *logger.Logger
	TokenSetting      *setting.TokenSettingS
)
