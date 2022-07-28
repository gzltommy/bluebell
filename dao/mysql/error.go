package mysql

import "errors"

var (
	ErrorUserExited    = errors.New("user exited")
	ErrorUserNotExit   = errors.New("user not exit")
	ErrorPasswordWrong = errors.New("password wrong")
)
