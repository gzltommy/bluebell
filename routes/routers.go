package routes

import (
	"bluebell/controller"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	r := initEngine(mode)
	v1 := r.Group("/api/v1")
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtAuthMiddleware())
	{
		// 用户
		v1.POST("/user/signup", controller.SignupHandler)
		v1.POST("/user/login", controller.LoginHandler)
		auth.GET("/user/refresh-token", controller.RefreshTokenHandler)
	}

	return r
}
