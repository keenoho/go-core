package service

import (
	"context"
	"fmt"
)

type MicroServiceServerInterface interface {
	UnimplementedServiceHandlerServer
	Send(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error)
}

type MicroServiceServer struct {
	UnimplementedServiceHandlerServer
	Service *MicroService
}

func (s *MicroServiceServer) Send(ctx context.Context, in *ServiceRequest) (*ServiceResponse, error) {
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
			return
		}
	}()

	// create context
	serviceContext := MicroServiceContext{
		ConnectContext: &ctx,
		RequestIn:      in,
		ContextData:    map[string]any{},
	}

	// call middlewares
	if len(s.Service.MiddlewareList) > 0 {
		for _, middleware := range s.Service.MiddlewareList {
			middleware(&serviceContext)
		}
	}

	// call handler
	res := handler(&serviceContext)

	return &ServiceResponse{
		Data: res.Data,
		Code: int64(res.Code),
		Msg:  res.Msg,
	}, nil
}
