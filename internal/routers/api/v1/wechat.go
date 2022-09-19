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

	param := validate.GetFriendListRequest{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	list := service.GetFriendList(param.RobotWxid, param.IsRefresh)
	response := app.NewResponse(c)
	response.ToResponse(list)
}

//GetGroupList 获取群聊列表
func (w Wechat) GetGroupList(c *gin.Context) {

	param := validate.GetFriendListRequest{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	list := service.GetGroupList(param.RobotWxid, param.IsRefresh)
	response := app.NewResponse(c)
	response.ToResponse(list)
}

//GetGroupMemberList 获取群成员列表
func (w Wechat) GetGroupMemberList(c *gin.Context) {
	param := validate.GetGroupMemberList{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	list := service.GetGroupMemberList(param.RobotWxid, param.GroupWxid, param.IsRefresh)
	response := app.NewResponse(c)
	response.ToResponse(list)
}

//GetGroupMemberList 获取群成员
func (w Wechat) GetGroupMember(c *gin.Context) {
	param := validate.GetGroupMember{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	//list :=
	service.GetGroupMember(param.RobotWxid, param.GroupWxid, param.MemberWxid)
	//response := app.NewResponse(c)
	//response.ToResponse(list)
}

// SendTextMsg  发送文本消息
func (w Wechat) SendTextMsg(c *gin.Context) {
	param := validate.SendMsgRequest{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	fmt.Println(param)
	service.SendTextMsg(param.RobotWxid, param.ToWxid, param.Msg)
	response := app.NewResponse(c)
	response.ToResponse(nil)
}

func (w Wechat) SendGroupAtMsg(c *gin.Context) {
	param := validate.SendMsgRequest{}
	_, _ = app.BindAndValid(c, &param)
	service := lovely_cat.NewWechatService(c)
	fmt.Println(param)
	service.SendGroupAtMsg(param.RobotWxid, param.GroupWxid, param.AtWxid, param.AtName, param.Msg)
	response := app.NewResponse(c)
	response.ToResponse(nil)
}
