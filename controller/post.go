package controller

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

func PostListHandler(c *gin.Context) {
	req := &model.ParamPostList{}
	if err := c.ShouldBindQuery(req); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
	}
	// 获取数据
	data, err := logic.GetPostList(req)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.GetPostList(%+v)", req))
		return
	}
	render.ResponseSuccess(c, data)
}

// PostList2Handler 升级版帖子列表接口
// @summary 升级版帖子列表接口
// @description 可按社区按时间或分数排序查询帖子列表接口
// @tags 帖子相关接口
// @accept application/json
// @produce application/json
// @param Authorization header string false "Bearer 用户令牌"
// @param object query model.ParamPostList2 false "查询参数"
// @security ApiKeyAuth
// @success 200 {object} _ResponsePostList
// @router /posts2 [get]
func PostList2Handler(c *gin.Context) {
	// GET 请求参数(query string)： /api/v1/posts2?page=1&size=10&order=time
	// 获取分页参数
	p := &model.ParamPostList2{
		Page:  1,
		Size:  10,
		Order: model.OrderTime, // magic string
	}
	if err := c.ShouldBindQuery(p); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
		return
	}

	// 获取数据
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.GetPostListNew(%+v)", p))
		return
	}
	render.ResponseSuccess(c, data)
}

func CreatePostHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var req = &model.ParamCreatePost{}
	if err := c.ShouldBindJSON(req); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
	}

	userId, err := getAuthUserID(c)
	if err != nil {
		render.ResponseError(c, render.CodeNotLogin, errors.WithMessage(err, "getAuthUserID fail"))
	}
	req.AuthorId = userId

	// 2、创建帖子
	err = logic.CreatePost(req)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.CreatePost(%+v)", err))
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
		render.ResponseError(c, render.CodeErrParams, errors.Wrapf(err, "strconv.ParseInt(%s, 10, 64)", postIdStr))
	}

	// 2、根据 id 取出 id 帖子数据(查数据库)
	post, err := logic.GetPostById(postId)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.GetPostById(%v)", postId))
	}

	// 3、返回响应
	render.ResponseSuccess(c, post)
}
