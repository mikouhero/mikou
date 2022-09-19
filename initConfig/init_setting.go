package initConfig

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"mikou/global"
	"time"
)

//  Setting  viper库读取后的实例
type Setting struct {
	vp *viper.Viper
}

// GDSetting  Setting 结构体的实例 别名
var GDSetting *Setting

//  Initialization 统一初始化
func Initialization() {

	e := initSetting()
	if e != nil {
		log.Fatalf("init.setting  err: %v", e)
		global.LoggerV2.Fatalf("init.setting  err: %v", e)
	}

	// 初始化db 配置
	initDb()

	// 初始化日志配置
	_ = initLog()
	initlog2()

	//初始化job 信息
	initCron()
}

// initSetting  启动时 获取config 配置信息
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
	err = GDSetting.ReadSection("Tencent", &global.TencentSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

// NewSetting 通过viper库 读取configs/config.yaml 文件的配置信息
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

// ReadSection 获取配置的节点信息
func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
