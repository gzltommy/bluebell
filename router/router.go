package router

import (
	"bluebell/controller"
	_ "bluebell/docs"
	"bluebell/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := initEngine()

	// 前端页面文件
	//r.LoadHTMLFiles("templates/index.html") // 加载 html
	//r.Static("/static", "./static")         // 加载静态文件
	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", nil)
	//})

	// 接口
	v1 := r.Group("/api/v1")
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtAuth())
	{
		v1.GET("/test", controller.Test)
	}
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
