package initConfig

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"mikou/global"
	"time"
)

func init() {
	Initialization()
}

// 统一初始化
func Initialization() {

	e := initSetting()
	if e != nil {
		log.Fatalf("init.setting  err: %v", e)
		global.LoggerV2.Fatalf("init.setting  err: %v", e)
	}

	initDb()

	_ = initLog()
	initlog2()
}

func initSetting() error {
	var err error
	GDSetting, err = NewSetting()
	if err != nil {
		return err
	}

	err = GDSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = GDSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = GDSetting.ReadSection("DatabaseV2", &global.DatabaseSettingV2)
	if err != nil {
		return err
	}

	err = GDSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = GDSetting.ReadSection("Token", &global.TokenSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

type Setting struct {
	vp *viper.Viper
}

var GDSetting *Setting

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		//Initialization()
		//todo
	})
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
