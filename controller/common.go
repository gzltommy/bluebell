package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ContextUserIDKey = "user_id"
)

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
