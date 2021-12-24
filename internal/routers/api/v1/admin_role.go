package v1

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type AdminRole struct {
}

func NewAdminRole() AdminRole {
	return AdminRole{}
}

// @Summary 获取用户列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.AdminRole "成功"
// @Failure 400 {object} code.Error "请求错误"
// @Failure 500 {object} code.Error "内部错误"
// @Router /api/v1/tags [get]
func (a AdminRole) List(c *gin.Context) {
	// 请求参数验证
	param := validate.ListAdminRoleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	//  service 实例
	svc := service.New(c.Request.Context())
	// 分页实例
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//  统计用户数量
	totalRows, err := svc.CountAdminRole(&param)

	if err != nil {
		global.Logger.Errorf("svc.CountAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	// service 数据结果
	list, err := svc.ListAdminRole(&param, &pager)

	if err != nil {
		global.Logger.Errorf("svc.ListAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponseList(list, totalRows)
	return
}

func (a AdminRole) Create(c *gin.Context) {
	// 请求参数验证
	param := validate.CreateAdminRoleRequest{}
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

	roleId, err := svc.CreateAdminRole(&param)

	if err != nil {
		global.Logger.Errorf("svc.CreateAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"role_id": roleId,
	}
	response.ToResponse(r)
	return
}

func (a AdminRole) Update(c *gin.Context) {
	param := validate.UpdateAdminRoleRequest{}
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
	err := svc.UpdateAdminRole(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a AdminRole) Delete(c *gin.Context) {
	param := validate.DeleteAdminRoleRequest{}
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
	err := svc.DeleteAdminRole(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a AdminRole) All(c *gin.Context) {

	response := app.NewResponse(c)

	//  service 实例
	svc := service.New(c.Request.Context())

	// service 数据结果
	list, err := svc.AllAdminRole()

	if err != nil {
		global.Logger.Errorf("svc.AllAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"list": list,
	}
	response.ToResponse(r)
	return
}


func (a AdminRole) Find(c *gin.Context) {
	param := validate.FindAdminRoleRequest{}
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
	role ,err := svc.FindAdminRole(&param)
	if err != nil {
		global.Logger.Errorf("svc.FindAdminRole err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"info": role,
	}
	response.ToResponse(r)
	return
}
