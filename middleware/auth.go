package middleware

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"bluebell/render"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带 Token 有三种方式 1.放在请求头 2.放在请求体 3.放在 URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			render.ResponseAbort(c, http.StatusUnauthorized, "Authorization header not provided")
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			render.ResponseAbort(c, http.StatusUnauthorized, "Authorization Bearer is error")
			return
		}

		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			render.ResponseAbort(c, http.StatusUnauthorized, "Authorization Bearer is error")
			return
		}

		// 将当前请求的 userID 信息保存到请求的上下文 c 上
		c.Set(controller.ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过 c.Get(controller.ContextUserIDKey) 来获取当前请求的用户信息
	}
}
