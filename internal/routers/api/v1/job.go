package v1

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	validate "mikou/internal/validate/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

type Job struct {
}

func NewJob() Job {
	return Job{}
}

// @Summary 获取Job列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Job "成功"
// @Failure 400 {object} code.Error "请求错误"
// @Failure 500 {object} code.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Job) List(c *gin.Context) {
	// 请求参数验证
	param := validate.ListJobRequest{}
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
	// 分页实例
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	//  统计Job数量
	totalRows, err := svc.CountJob(&param)

	if err != nil {
		//global.Logger.Errorf("svc.CountJob err: %v", err)
		global.LoggerV2.Errorf("app.CountJob errs: %v", err)

		response.ToErrorResponse(errcode.ServerError)
		return
	}

	// service 数据结果
	list, err := svc.ListJob(&param, &pager)

	if err != nil {
		//global.Logger.Errorf("svc.ListJob err: %v", err)
		global.LoggerV2.Error("app.ListJob errs", err)

		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponseList(list, totalRows)
	return
}

func (a Job) Create(c *gin.Context) {
	// 请求参数验证
	param := validate.CreateJobRequest{}
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

	Id, err := svc.CreateJob(&param)

	if err != nil {
		//global.Logger.Errorf("svc.CreateJob err: %v", err)
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	var r = map[string]interface{}{
		"id": Id,
	}
	response.ToResponse(r)
	return
}

func (a Job) Update(c *gin.Context) {
	param := validate.UpdateJobRequest{}
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

	err := svc.UpdateJob(&param)
	if err != nil {
		//global.Logger.Errorf("svc.UpdateJob err: %v", err)
		global.LoggerV2.Errorf("app.BindAndValid errs: %v", errs)

		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}

func (a Job) Delete(c *gin.Context) {
	param := validate.DeleteJobRequest{}
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
	err := svc.DeleteJob(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteJob err: %v", err)
		response.ToErrorResponse(errcode.ServerError)
		return
	}
	response.ToResponse(nil)
	return
}
