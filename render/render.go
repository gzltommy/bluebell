package render

import (
	"bluebell/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
)

type RespJsonData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"` // omitempty 当 data 为空时,不展示这个字段
}

func ResponseAbort(c *gin.Context, code int, err error) {
	result := &RespJsonData{
		Code: code,
		Msg:  errorMsg(c, code, err),
		Data: nil,
	}
	c.AbortWithStatusJSON(code, result)
}

func ResponseError(c *gin.Context, code int, err error) {
	result := &RespJsonData{
		Code: code,
		Msg:  errorMsg(c, code, err),
		Data: nil,
	}
	c.JSON(http.StatusOK, result)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	result := &RespJsonData{
		Code: CodeSucceed,
		Msg:  code2Msg(CodeSucceed),
		Data: data,
	}
	c.JSON(http.StatusOK, result)
}

func errorMsg(c *gin.Context, code int, err error) string {
	eCode := utils.Err2Hash(err)
	httpRequest, _ := httputil.DumpRequest(c.Request, false)
	zap.L().Error("[errorCode]",
		zap.Int("response code", code),
		zap.String("original error", fmt.Sprintf("%T,%v", errors.Cause(err), errors.Cause(err))),
		zap.String("error stack trace", fmt.Sprintf("%+v", err)), // %+v 打印 error 堆栈信息
		zap.String("error code", eCode),
		zap.String("request", string(httpRequest)),
		zap.String("stack", string(debug.Stack())),
	)
	//return code2Msg(code) + fmt.Sprintf("error_code:%s", eCode)
	return code2Msg(code) + fmt.Sprintf("error:%v", err)
}

func code2Msg(code int) string {
	if v, ok := codeMsg[code]; ok {
		return v
	} else {
		return codeMsg[CodeServerBusy]
	}
}
