package core

import (
	"github.com/gin-gonic/gin"
)

type ControllerFunc func(ctx *gin.Context) (ResponseData, int)

type ErrorData struct {
	Error  any    `json:"error"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

type ResponseData struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SysTime int64  `json:"systime"`
}

type SignatureData struct {
	Expired int    `json:"expired"`
	App     int64  `json:"app"`
	Sig     string `json:"sig"`
	Ttl     int    `json:"ttl"`
	Ts      int64  `json:"ts"`
	Nonce   string `json:"nonce"`
	Data    string `json:"data"`
}

type SignatureInnerData struct {
	Id      int64   `json:"id"`
	Account string  `json:"account"`
	Name    string  `json:"name"`
	Role    []int64 `json:"role"`
}
