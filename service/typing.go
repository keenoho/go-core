package service

type ResponseData struct {
	Data []byte `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
