package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/model"
	"bluebell/pkg/jwt"
	"bluebell/render"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strings"
)

func SignupHandler(c *gin.Context) {
	// 1.参数校验
	req := &model.ParamSignUp{}
	if err := c.ShouldBindJSON(req); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
	}

	// 2.逻辑处理
	err := logic.SignUp(req)
	if err != nil {
		if errors.Is(err, mysql.ErrorUserExited) {
			render.ResponseError(c, render.CodeUserExisted, errors.WithMessagef(err, "logic.SignUp(%+v)", req))
			return
		}
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "logic.SignUp(%+v)", req))
		return
	}

	// 3.返回数据
	render.ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1.参数校验
	req := &model.ParamLogin{}
	if err := c.ShouldBindJSON(req); err != nil {
		render.ResponseError(c, render.CodeErrParams, errors.WithStack(err))
	}

	// 2.逻辑处理
	user, err := logic.Login(req)
	if err != nil {
		switch {
		case errors.Is(err, mysql.ErrorUserNotExit):
			render.ResponseError(c, render.CodeUserNotExit, err)
			return
		case errors.Is(err, mysql.ErrorPasswordWrong):
			render.ResponseError(c, render.CodePasswordWrong, err)
			return
		}
		render.ResponseError(c, render.CodeServerBusy, err)
		return
	}

	// 3.返回数据
	render.ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID), // js 识别的最大值：id 值大于1<<53-1  int64: i<<63-1
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带 Token 有三种方式 1.放在请求头 2.放在请求体 3.放在 URI
	// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		render.ResponseError(c, render.CodeInvalidToken, errors.New("请求头缺少 Auth Token"))
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		render.ResponseError(c, render.CodeInvalidToken, errors.New("Token 格式不对"))
		return
	}

	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		render.ResponseError(c, render.CodeServerBusy, errors.WithMessagef(err, "jwt.RefreshToken(%v, %v)", parts[1], rt))
		return
	}

	render.ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
