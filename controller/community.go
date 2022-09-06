package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

// CommunityListHandler 社区列表
// @Summary 社区列表
// @Description 社区列表
// @Tags 社区业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query model.Community false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community [get]
func CommunityListHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name)以列表的形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessage(err, "logic.GetCommunityList() fail"))
		return
	}
	render.ResponseSuccess(c, communityList)
}

// CommunityDetailHandler 社区详情
// @Summary 社区详情
// @Description 社区详情
// @Tags 社区业务接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query communityId     path    int     true        "id"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1、获取社区ID
	communityIdStr := c.Param("id")                               // 获取URL参数
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64) // id字符串格式转换
	if err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.Wrapf(err, "strconv.ParseUint(%s, 10, 64)", communityIdStr))
		return
	}

	// 2、根据ID获取社区详情
	communityList, err := logic.GetCommunityDetailByID(communityId)
	if err != nil {
		switch {
		case errors.Is(err, mysql.ErrorInvalidID):
			render.ResponseError(c, render.CodeInvalidCommunityID, err)
			return
		case errors.Is(err, mysql.ErrorQueryFailed):
			render.ResponseError(c, render.CodeServerError, err)
			return
		}
		render.ResponseError(c, render.CodeServerBusy, err)
		return
	}
	render.ResponseSuccess(c, communityList)
}
