package service

import (
	"context"
	"encoding/json"
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
	Mode    string
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
		return nil, fmt.Errorf("handler is not in routeMap: %s", in.Url)
	}
	defer func() {
		err := recover()
		if err != nil {
			s.PrintDebug("handler send error: %v", err)
			return
		}
	}()
	res := handler(ctx, in)
	var resByte []byte
	if res.Data != nil {
		resByte, _ = json.Marshal(res.Data)
	}

	return &ServiceMsgResponse{
		Data: resByte,
		Code: int64(res.Code),
		Msg:  res.Msg,
	}, nil
}
