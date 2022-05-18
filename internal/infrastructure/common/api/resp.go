package api

import (
	"github.com/gin-gonic/gin"
	"go/types"
	"net/http"
)

const (
	SUCCESS = "0"
	FAILED  = "-1"
	OK      = "ok"
)

type Result[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// Response OK
func RespOk[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, &Result[T]{Code: SUCCESS, Msg: OK, Data: data})
}

// Response Error
func RespError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Result[types.Nil]{Code: FAILED, Msg: msg})
}

// Response Error
func RespErrorWithCode(c *gin.Context, code string, msg string) {
	c.JSON(http.StatusOK, &Result[types.Nil]{Code: code, Msg: msg})
}
