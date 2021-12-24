package app

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/pkg/convert"
)

// 统一获取请求页数 以 GET 为主
func GetPage(c *gin.Context) int {
	var page int
	page = convert.StrTo(c.Query("page")).MustInt()

	if (page == 0) {
		page = convert.StrTo(c.PostForm("page")).MustInt()
	}

	if page <= 0 {
		return 1
	}

	return page
}

// 统一获取请求每页数量 以 GET 为主
func GetPageSize(c *gin.Context) int {
	var pageSize int

	pageSize = convert.StrTo(c.Query("page_size")).MustInt()

	if (pageSize == 0) {
		pageSize = convert.StrTo(c.PostForm("page_size")).MustInt()
	}

	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
