package main

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/initConfig"
	"mikou/internal/routers"
	"net/http"
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
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
