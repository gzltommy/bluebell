package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/render"
	"bluebell/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityListHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name)以列表的形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		render.ResponseError(c, render.CodeServerBusy)
		return
	}
	render.ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	// 1、获取社区ID
	communityIdStr := c.Param("id")                               // 获取URL参数
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64) // id字符串格式转换
	if err != nil {
		render.ResponseError(c, render.CodeErrParams)
		return
	}

	// 2、根据ID获取社区详情
	communityList, err := logic.GetCommunityDetailByID(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID() failed", zap.Error(err), zap.String("error code", utils.Err2Hash(err)))
		switch {
		case errors.Is(err, mysql.ErrorInvalidID):
			render.ResponseError(c, render.CodeInvalidCommunityID)
			return
		case errors.Is(err, mysql.ErrorQueryFailed):
			render.ResponseError(c, render.CodeServerError)
			return
		}
		render.ResponseError(c, render.CodeServerBusy, utils.Err2Hash(err))
		return
	}
	render.ResponseSuccess(c, communityList)
}
