package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mikou/global"
)

type Wechat struct {
}

func NewWechat() Wechat {
	return Wechat{}
}

// @Summary 获取Job列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Job "成功"
// @Failure 400 {object} code.Error "请求错误"
// @Failure 500 {object} code.Error "内部错误"
// @Router /api/v1/tags [get]
func (w Wechat) Message(c *gin.Context) {

	fmt.Println(c.Request)
	fmt.Println(c.Request.PostForm)

	global.LoggerV2.Infof("reveive %v", c.PostForm("type"))
	fmt.Println("reveive")
}
