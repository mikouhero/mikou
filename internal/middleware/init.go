package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {

	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// Secure is a middleware function that appends security
	r.Use(Secure)
	//r.Use(Recovery())
	r.Use(AccessLog())
}
