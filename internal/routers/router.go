package routers

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/internal/middleware"
	v1 "mikou/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	middleware.InitMiddleware(r)
	adminUser := v1.NewAdminUser()
	adminRole := v1.NewAdminRole()
	adminMenu := v1.NewAdminMenu()
	adminLogin := v1.NewAdminLogin()

	r.Any("/admin/login", adminLogin.Login)

	auth := r.Group(global.AppSetting.AuthPrefix)
	{
		auth.Use(middleware.JWT())
		auth.Any("/user/list", adminUser.List)
		auth.Any("/user/create", adminUser.Create)
		auth.Any("/user/update", adminUser.Update)
		auth.Any("/user/delete", adminUser.Delete)

		auth.Any("/role/list", adminRole.List)
		auth.Any("/role/all", adminRole.All)

		auth.Any("/role/create", adminRole.Create)
		auth.Any("/role/update", adminRole.Update)
		auth.Any("/role/delete", adminRole.Delete)
		auth.Any("/role/find", adminRole.Find)

		auth.Any("/menu/list", adminMenu.List)
		auth.Any("/menu/create", adminMenu.Create)
		auth.Any("/menu/update", adminMenu.Update)
		auth.Any("/menu/delete", adminMenu.Delete)
		auth.Any("/menu/find", adminMenu.Find)


		auth.Any("/login/logout", adminLogin.LogOut)
		auth.Any("/login/user_info", adminLogin.UserInfo)

	}
	return r
}
