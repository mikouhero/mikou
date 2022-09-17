package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/internal/service/lovely_cat"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type Wechat struct {
}

//NewWechat  接收微信的消息推送事件
func NewWechat() Wechat {
	return Wechat{}
}

// Message 统一接收微信的消息事件
func (w Wechat) Message(c *gin.Context) {

	// 请求参数验证
	param := validate.MessageRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	fmt.Println("reveive")

}

// GetFriendList 获取当前好友列表
func (w Wechat) GetFriendList(c *gin.Context) {
	service := lovely_cat.NewWechatService(c)
	list := service.GetFriendList("", 1)
	//list := service.GetLoggedAccountList()
	//list := service.GetGroupList("wxid_ysq8klcfelmf22", 0)
	service.GetGroupMemberList("wxid_ysq8klcfelmf22", "24976121095@chatroom", 1)
	response := app.NewResponse(c)
	response.ToResponse(list)
}
