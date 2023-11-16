package service

type ResponseData struct {
	Data []byte `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MicroServiceControllerFunc func(ctx *MicroServiceContext) ResponseData

type MicroServiceMiddleware func(ctx *MicroServiceContext)
