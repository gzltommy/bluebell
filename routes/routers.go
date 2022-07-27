package routes

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := initEngine()
	v1 := r.Group("/v1")
	v1.GET("/ping", controller.Pong)

	{
		// 用户
		v1.POST("user/signup", controller.Signup)
	}
	return r
}
