package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "user_id"
)

var (
	ErrorUserNotLogin = errors.New("user not login")
)

func getAuthUserID(c *gin.Context) (userID uint64, err error) {
	_userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = _userID.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
