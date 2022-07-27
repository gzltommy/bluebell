package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/render"
	"bluebell/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Signup(c *gin.Context) {
	// 1.参数校验
	req := models.ParamSignUp{}
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("logic.Signup with invalid param", zap.Error(err))
		render.Json(c, render.ErrParams, nil)
	}

	// 2.逻辑处理
	err := logic.SignUp(&req)
	if err != nil {
		if errors.Is(err, logic.ErrorUserExited) {
			render.Json(c, render.UserExisted, nil)
			return
		}
		zap.L().Error("logic.Signup failed", zap.Error(err), zap.String("error code", utils.Err2Hash(err)))
		render.Json(c, render.ErrorUndefined, nil, utils.Err2Hash(err))
		return
	}

	// 3.返回数据
	render.Json(c, render.Succeed, nil)
}
