package mysql

import "github.com/pkg/errors"

var (
	ErrorUserExited    = errors.New("user exited")
	ErrorUserNotExit   = errors.New("user not exit")
	ErrorPasswordWrong = errors.New("password wrong")
	ErrorInsertFailed  = errors.New("insert data failed")

	ErrorInvalidID   = errors.New("invalid id")
	ErrorQueryFailed = errors.New("query failed")
)
