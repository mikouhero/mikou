package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mikou/global"
	service "mikou/internal/service/v1"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
	"strings"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {

			_, claims, err := app.ParseToken(token)

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				if claims.UserAccount != global.AppSetting.SuperAdminUser {
					svc := service.New(c.Request.Context())
					menus, _ := svc.GetRoleMenuList(claims.RoleId)
					var authList []string
					for _, value := range menus {
						if value.Urls != "" && strings.HasPrefix(value.Urls, global.AppSetting.AuthPrefix) {
							authList = append(authList, strings.ToLower(value.Urls))
						}
					}
					if !app.InArrayString(authList, strings.ToLower(c.Request.URL.Path)) {
						ecode = errcode.Forbidden
					}
				}
				c.Set("UserId", claims.UserId)
				c.Set("UserAccount", claims.UserAccount)
				c.Set("UserRoleId", claims.RoleId)
			}

		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
