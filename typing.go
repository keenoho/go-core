package go_core

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
