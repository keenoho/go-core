package service

import (
	"context"
	"fmt"

	"github.com/keenoho/go-core"
)

type MicroServiceServerInterface interface {
	UnimplementedServiceMsgHandlerServer
	core.Logger
	Send(ctx context.Context, in *ServiceMsgRequest) (*ServiceMsgResponse, error)
}

type MicroServiceServer struct {
	UnimplementedServiceMsgHandlerServer
	core.Logger
	Service *MicroService
}

func (s *MicroServiceServer) Send(ctx context.Context, in *ServiceMsgRequest) (*ServiceMsgResponse, error) {
	if s.Service == nil {
		return nil, fmt.Errorf("service is not exist")
	}
	if s.Service.RouteMap == nil {
		return nil, fmt.Errorf("service routeMap is not exist")
	}
	if in.Url == "" || len(in.Url) < 1 {
		return nil, fmt.Errorf("income url is not exist")
	}
	handler, isExist := s.Service.RouteMap[in.Url]
	if !isExist {
		return nil, fmt.Errorf("handler is not in routeMap: %v", in.Url)
	}
	defer func() {
		err := recover()
		if err != nil {
			s.PrintError("handler send error: %v", err)
			return
		}
	}()

	// create context
	serviceContext := MicroServiceContext{
		ConnectContext: &ctx,
		RequestIn:      in,
	}

	// call middlewares
	if len(s.Service.MiddlewareList) > 0 {
		for _, middleware := range s.Service.MiddlewareList {
			middleware(&serviceContext)
		}
	}

	// call handler
	res := handler(&serviceContext)
	dataByte, _ := DataToBytes(res.Data)

	return &ServiceMsgResponse{
		Data: dataByte,
		Code: int64(res.Code),
		Msg:  res.Msg,
	}, nil
}
