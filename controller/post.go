package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
