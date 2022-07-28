package routes

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	r := initEngine(mode)
	v1 := r.Group("/v1")
	v1.GET("/ping", controller.Pong)

	{
		// 用户
		v1.POST("user/signup", controller.SignupHandler)
		v1.POST("user/login", controller.LoginHandler)
		v1.GET("user/refresh-token", controller.RefreshTokenHandler)
	}
	return r
}
