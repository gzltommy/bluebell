package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func PostListHandler(c *gin.Context) {
	req := &model.ParamPostList{}
	if err := c.ShouldBindQuery(req); err != nil {
		zap.L().Error("PostListHandler with invalid param", zap.Error(err))
		render.ResponseError(c, render.CodeErrParams)
	}
	// 获取数据
	data, err := logic.GetPostList(req)
	if err != nil {
		zap.L().Error("logic.GetPostList(req) fail", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, data)
}

func PostList2Handler(c *gin.Context) {
	// GET请求参数(query string)： /api/v1/posts2?page=1&size=10&order=time
	// 获取分页参数
	p := &model.ParamPostList2{
		Page:  1,
		Size:  10,
		Order: model.OrderTime, // magic string
	}
	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("PostList2Handler with invalid params", zap.Error(err))
		render.ResponseError(c, render.CodeErrParams)
		return
	}

	// 获取数据
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, data)
}

func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var req = &model.ParamCreatePost{}
	if err := c.ShouldBindJSON(req); err != nil {
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		render.ResponseError(c, render.CodeErrParams)
	}

	userId, err := getAuthUserID(c)
	if err != nil {
		zap.L().Error("LoginHandler getAuthUserID fail", zap.Error(err))
		render.ResponseError(c, render.CodeNotLogin)
	}
	req.AuthorId = userId

	// 2、创建帖子
	err = logic.CreatePost(req)
	if err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
		return
	}

	// 3、返回响应
	render.ResponseSuccess(c, nil)
}

func PostDetailHandler(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("PostDetailHandler  strconv.ParseInt(postIdStr,10,64) fail", zap.Error(err))
		render.ResponseError(c, render.CodeErrParams)
	}

	// 2、根据 id 取出 id 帖子数据(查数据库)
	post, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
	}

	// 3、返回响应
	render.ResponseSuccess(c, post)
}
