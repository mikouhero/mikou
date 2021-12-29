package v1

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type AdminLogin struct{}

func NewAdminLogin() AdminLogin {
	return AdminLogin{}
}

func (a *AdminLogin) Login(c *gin.Context) {
	param := validate.LoginAdminLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	s, e := svc.AdminUserLogin(param)
	if e != nil {
		response.ToErrorResponse(errcode.ServerError.WithDetails(e.Error()))
		return
	}

	var r = map[string]interface{}{
		"token": s,
	}
	response.ToResponse(r)
	return

}
func (a *AdminLogin) UserInfo(c *gin.Context) {
	response := app.NewResponse(c)
	userIdInterface, exists := c.Get("UserId")
	if !exists {
		response.ToErrorResponse(errcode.ServerError)
	}
	svc := service.New(c.Request.Context())
	roleIdInterface, exists := c.Get("UserRoleId")
	if !exists {
		response.ToErrorResponse(errcode.ServerError)
	}
	var roleId int
	var userId int
	switch roleIdInterface.(type) {
	case int:
		roleId = roleIdInterface.(int)
	default:
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	switch userIdInterface.(type) {
	case int:
		userId = userIdInterface.(int)
	default:
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	lists, _ := svc.GetRoleMenuTreeList(roleId)
	user := svc.GetUserByID(userId)

	var r = map[string]interface{}{
		"userinfo": user,
		"menulist": lists,
	}
	response.ToResponse(r)
	return
}

func (a *AdminLogin) LogOut(c *gin.Context) {
	response := app.NewResponse(c)
	// todo
	response.ToResponse("")
	return
}
