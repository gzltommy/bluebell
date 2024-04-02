package controller

import (
	model2 "bluebell/dao/mysql/op"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// CreateCommentHandler 创建评论
// @summary 创建评论
// @description 创建评论接口
// @tags 评论
// @accept application/json
// @produce application/json
// @param Authorization header string true "Bearer Token"
// @param object query  model.Comment false "create param"
// @security ApiKeyAuth
// @success 200 {object} _RespComment
// @router /comment [post]
func CreateCommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
		return
	}

	// 获取作者 ID，当前请求的 UserID
	userID, err := getAuthUserID(c)
	if err != nil {
		render.ResponseError(c, render.CodeNotLogin, errors.WithMessage(err, "getAuthUserID fail"))
		return
	}
	comment.CommentID = uint64(snowflake.GenID()) // 生成评论ID
	comment.AuthorID = userID

	// 创建评论
	if err := model2.CreateComment(&comment); err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "CreateComment(%+v)", comment))
		return
	}
	render.ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		render.ResponseError(c, render.CodeErrParams, errors.New("ids parameter not found"))
		return
	}
	posts, err := model2.GetCommentListByIDs(ids)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "GetCommentListByIDs(%v) fail", ids))
		return
	}
	render.ResponseSuccess(c, posts)
}
