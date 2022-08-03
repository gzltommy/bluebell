package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(c *gin.Context) {
	// 参数校验,给哪个文章投什么票
	vote := new(model.ParamVote)
	if err := c.ShouldBindJSON(&vote); err != nil {
		zap.L().Error("PostListHandler with invalid param", zap.Error(err))
		render.ResponseError(c, render.CodeErrParams)
		return
	}
	// 获取当前请求用户的id
	userID, err := getAuthUserID(c)
	if err != nil {
		render.ResponseError(c, render.CodeNotLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, vote); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, nil)
}
