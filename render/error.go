package render

import "fmt"

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg ...string) *Error {
	return &Error{
		Code: 0,
		Msg:  code2Msg(code, msg...),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d,message:%s", e.Code, e.Msg)
}
