package v1

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type AdminMenu struct {
}

func NewAdminMenu() AdminMenu {
	return AdminMenu{}
}

// @Summary 获取用户列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.AdminMenu "成功"
// @Failure 400 {object} code.Error "请求错误"
// @Failure 500 {object} code.Error "内部错误"
// @Router /api/v1/tags [get]
func (a AdminMenu) List(c *gin.Context) {
	svc := service.New(c.Request.Context())
	list, _ := svc.AllAdminMenu()
	response := app.NewResponse(c)
	response.ToResponse(list)

}

func (a AdminMenu) Create(c *gin.Context) {
	// 请求参数验证
	param := validate.CreateAdminMenuRequest{}
	//  service 实例
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		//global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)

		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	//  service 实例
	svc := service.New(c.Request.Context())

	roleId, err := svc.CreateAdminMenu(&param)

	if err != nil {
		global.LoggerV2.Errorf("svc.CreateAdminMenu err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"role_id": roleId,
	}
	response.ToResponse(r)
	return
}

func (a AdminMenu) Update(c *gin.Context) {
	param := validate.UpdateAdminMenuRequest{}
	//  service 实例
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//  service 实例
	svc := service.New(c.Request.Context())
	err := svc.UpdateAdminMenu(&param)
	if err != nil {
		global.LoggerV2.Errorf("svc.UpdateAdminMenu err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a AdminMenu) Delete(c *gin.Context) {
	param := validate.DeleteAdminMenuRequest{}
	//  service 实例
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//  service 实例
	svc := service.New(c.Request.Context())
	err := svc.DeleteAdminMenu(&param)
	if err != nil {
		global.LoggerV2.Errorf("svc.DeleteAdminMenu err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a AdminMenu) Find(c *gin.Context) {
	param := validate.FindAdminMenuRequest{}
	//  service 实例
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//  service 实例
	svc := service.New(c.Request.Context())
	menu, e := svc.FindAdminMenu(&param)
	err := e
	if err != nil {
		global.LoggerV2.Errorf("svc.DeleteAdminMenu err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"info": menu,
	}
	response.ToResponse(r)
	return
}

