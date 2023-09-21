package service

import "context"

type MicroServiceContextInterface interface {
}

type MicroServiceContext struct {
	ConnectContext *context.Context
	RequestIn      *ServiceMsgRequest
	Data           map[string]any
}

func (ctx *MicroServiceContext) Get(key string) any {
	return ctx.Data[key]
}

func (ctx *MicroServiceContext) Set(key string, value any) {
	ctx.Data[key] = value
}
