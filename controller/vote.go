package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func VoteHandler(c *gin.Context) {
	// 参数校验,给哪个文章投什么票
	vote := new(model.ParamVote)
	if err := c.ShouldBindJSON(&vote); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
		return
	}
	// 获取当前请求用户的id
	userID, err := getAuthUserID(c)
	if err != nil {
		render.ResponseError(c, render.CodeNotLogin, errors.WithMessage(err, "getAuthUserID fail"))
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, vote); err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.VoteForPost(%v,%+v)", userID, vote))
		return
	}
	render.ResponseSuccess(c, nil)
}
