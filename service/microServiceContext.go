package service

import "context"

type MicroServiceContextInterface interface {
}

type MicroServiceContext struct {
	ConnectContext *context.Context
	RequestIn      *ServiceRequest
	ContextData    map[string]any
}

func (ctx *MicroServiceContext) Get(key string) any {
	return ctx.ContextData[key]
}

func (ctx *MicroServiceContext) Set(key string, value any) {
	ctx.ContextData[key] = value
}
