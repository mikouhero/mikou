package middleware

import (
	"github.com/gin-gonic/gin"
	"mikou/global"
	"mikou/pkg/app"
	"mikou/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				global.LoggerV2.Errorf("panic recover err: %v", err)

				// todo 报警

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}