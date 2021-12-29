package initConfig

import (
	"mikou/global"

	model "mikou/internal/model/v1"
	modelv2 "mikou/internal/model/v2"
)

func initDb() {
	e := setupDBEngineV2()
	if e!=nil {
		//log.Fatalf("init.db err: %v", e)
		global.LoggerV2.Errorf("init.db err: %v", e)
	}
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngineV2() error {
	var err error
	global.DBEngineV2, err = modelv2.NewDBEngine(global.DatabaseSettingV2)
	if err != nil {
		return err
	}
	return nil
}
