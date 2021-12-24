package v1

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type AdminUser struct {
}

func NewAdminUser() AdminUser {
	return AdminUser{}
}

// @Summary 获取用户列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.AdminUser "成功"
// @Failure 400 {object} code.Error "请求错误"
// @Failure 500 {object} code.Error "内部错误"
// @Router /api/v1/tags [get]
func (a AdminUser) List(c *gin.Context) {
	// 请求参数验证
	param := validate.ListAdminUserRequest{}
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
	totalRows, err := svc.CountAdminUser(&param)


	if err != nil {
		global.Logger.Errorf("svc.CountAdminUser err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	// service 数据结果
	list, err := svc.ListAdminUser(&param, &pager)

	if err != nil {
		global.Logger.Errorf("svc.ListAdminUser err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponseList(list, totalRows)
	return
}

func (a AdminUser) Create(c *gin.Context) {
	// 请求参数验证
	param := validate.CreateAdminUserRequest{}
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
	exist := svc.CheckAccountExist(param.Account)
	if exist != nil {
		response.ToErrorResponse(errcode.AdminUserAccountHadExist)
		return
	}

	userId, err := svc.CreateAdminUser(&param)

	if err != nil {
		global.Logger.Errorf("svc.CreateAdminUser err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"user_id": userId,
	}
	response.ToResponse(r)
	return
}

func (a AdminUser) Update(c *gin.Context) {
	param := validate.UpdateAdminUserRequest{}
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
	if  param.Account != "" {
		exist := svc.CheckAccountExist(param.Account)
		if exist != nil && exist.ID != param.Id {
			response.ToErrorResponse(errcode.AdminUserAccountHadExist)
			return
		}
	}


	err := svc.UpdateAdminUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateAdminUser err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a AdminUser) Delete(c *gin.Context) {
	param := validate.DeleteAdminUserRequest{}
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
	err := svc.DeleteAdminUser(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteAdminUser err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}
