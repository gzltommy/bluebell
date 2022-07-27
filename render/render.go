package render

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespJsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"` // omitempty 当 data 为空时,不展示这个字段
}

func AbortJson(c *gin.Context, code int, data interface{}, errHashV ...string) {
	result := &RespJsonData{
		Code: code,
		Msg:  getMessage(code, errHashV...),
		Data: data,
	}
	c.AbortWithStatusJSON(code, result)
}

func Json(c *gin.Context, code int, data interface{}, errHashV ...string) {
	result := &RespJsonData{
		Code: code,
		Msg:  getMessage(code, errHashV...),
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

func getMessage(code int, errHashV ...string) string {
	if code == ErrorUndefined && len(errHashV) > 0 {
		return fmt.Sprintf("undefined error.error code:%s", errHashV[0])
	}
	return codeMsg[code]
}
