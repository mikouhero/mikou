package initConfig

import (
	"mikou/global"

	model "mikou/internal/model/v1"
	modelv2 "mikou/internal/model/v2"
)

// initDb  初始化数据库配置
func initDb() {
	e := setupDBEngineV2()
	if e != nil {
		//log.Fatalf("init.db err: %v", e)
		global.LoggerV2.Errorf("init.db err: %v", e)
	}
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

// setupDBEngineV2 采用新版的gorm  初始化数据库配置
func setupDBEngineV2() error {
	var err error
	global.DBEngineV2, err = modelv2.NewDBEngine(global.DatabaseSettingV2)

	return err
}
