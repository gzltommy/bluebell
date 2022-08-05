package controller

import (
	"bluebell/dao/mysql"
	"bluebell/model"
	"bluebell/pkg/snowflake"
	"bluebell/render"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCommentHandler 创建评论
func CreateCommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		fmt.Println(err)
		render.ResponseError(c, render.CodeErrParams)
		return
	}

	// 获取作者 ID，当前请求的 UserID
	userID, err := getAuthUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		render.ResponseError(c, render.CodeNotLogin)
		return
	}
	comment.CommentID = uint64(snowflake.GenID()) // 生成评论ID
	comment.AuthorID = userID

	// 创建评论
	if err := mysql.CreateComment(&comment); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		render.ResponseError(c, render.CodeErrParams)
		return
	}
	posts, err := mysql.GetCommentListByIDs(ids)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, posts)
}
