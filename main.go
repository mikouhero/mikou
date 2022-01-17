package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/initConfig"
	"mikou/internal/routers"
	"mikou/pkg/app"
	"syscall"
)

func init() {
	initConfig.Initialization()
}

// @title mikou PIM 管理系统
// @version 1.0
// @description 接口文档
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	endless.DefaultReadTimeOut = global.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = global.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	s := endless.NewServer(":"+global.ServerSetting.HttpPort, router)
	s.BeforeBegin = func(add string) {
		global.LoggerV2.Infof("System begin  and  Actual pid is %d", syscall.Getpid())
	}

	//s := &http.Server{
	//	Addr:           ":" + global.ServerSetting.HttpPort,
	//	Handler:        router,
	//	ReadTimeout:    global.ServerSetting.ReadTimeout,
	//	WriteTimeout:   global.ServerSetting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	_ = s.ListenAndServe()

}
