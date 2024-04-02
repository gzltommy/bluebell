package controller

import (
	"bluebell/render"
	"github.com/gin-gonic/gin"
)

type reqAiCompanionDetail struct {
	AiCompanionID int64  `form:"ai_companion_id" json:"ai_companion_id" binding:"required"`
	Platform      string `form:"platform" binding:"omitempty,oneof=web win mac android ios"`
}

func Test(c *gin.Context) {
	var req reqAiCompanionDetail
	if err := c.ShouldBindQuery(&req); err != nil {
		render.ResponseError(c, render.CodeErrParams, err)
		return
	}
	render.ResponseSuccess(c, req)
}
