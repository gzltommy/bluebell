package router

import (
	"bluebell/controller"
	_ "bluebell/docs"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"

	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetupRouter(mode string) *gin.Engine {
	r := initEngine(mode)

	// 注册 swagger
	//r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
	r.GET("/swagger/*any", gs.DisablingWrapHandler(sf.Handler, "NAME_OF_ENV_VARIABLE"))

	v1 := r.Group("/api/v1")
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtAuthMiddleware())
	{
		v1.POST("/user/signup", controller.SignupHandler)
		v1.POST("/user/login", controller.LoginHandler)
		auth.GET("/user/refresh-token", controller.RefreshTokenHandler)
	}
	{
		v1.GET("/community", controller.CommunityListHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}
	{
		auth.POST("/post/create", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.PostDetailHandler)  // 查询帖子详情
		v1.GET("/post/list", controller.PostListHandler)   // 分页展示帖子列表
		v1.GET("/post/list2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表
	}
	{
		auth.POST("/vote", controller.VoteHandler) // 投票
	}
	{
		auth.POST("/comment", controller.CreateCommentHandler)
		v1.GET("/comment", controller.CommentListHandler)
	}

	return r
}
