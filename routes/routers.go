package routes

import (
	"bluebell/controller"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	r := initEngine(mode)
	v1 := r.Group("/v1")
	auth := r.Group("/v1")
	auth.Use(middleware.JwtAuthMiddleware())
	{
		// 用户
		v1.POST("user/signup", controller.SignupHandler)
		v1.POST("user/login", controller.LoginHandler)
		v1.GET("user/refresh-token", controller.RefreshTokenHandler)
	}

	return r
}
