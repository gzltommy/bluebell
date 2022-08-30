package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespJsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty 当 data 为空时,不展示这个字段
}

func ResponseAbort(c *gin.Context, code int, data interface{}) {
	result := &RespJsonData{
		Code: code,
		Msg:  http.StatusText(code),
		Data: data,
	}
	c.AbortWithStatusJSON(code, result)
}

func ResponseError(c *gin.Context, code int, msg ...string) {
	result := &RespJsonData{
		Code: code,
		Msg:  getCodeMsg(code, msg...),
		Data: nil,
	}
	c.JSON(http.StatusOK, result)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	result := &RespJsonData{
		Code: CodeSucceed,
		Msg:  getCodeMsg(CodeSucceed),
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

func getCodeMsg(code int, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	if v, ok := codeMsg[code]; ok {
		return v
	} else {
		return codeMsg[CodeServerBusy]
	}
}
