package core

type ErrorData struct {
	Error  any    `json:"error"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

type ResponseData struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

type ServiceResponseData struct {
	Data any    `json:"data"`
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
