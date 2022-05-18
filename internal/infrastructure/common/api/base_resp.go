package api

import "go/types"

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
func RespOk[T any](data T) *Result[T] {
	return &Result[T]{Code: SUCCESS, Msg: OK, Data: data}
}

// Response Error
func RespError(msg string) *Result[types.Nil] {
	return &Result[types.Nil]{Code: FAILED, Msg: msg}
}

// Response Error
func RespErrorWithCode(code string, msg string) *Result[types.Nil] {
	return &Result[types.Nil]{Code: code, Msg: msg}
}
