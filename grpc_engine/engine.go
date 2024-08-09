package grpc_engine

import (
	"net"

	"google.golang.org/grpc"
)

type RegisterServiceOption struct {
	ServiceDesc   *grpc.ServiceDesc
	ServiceServer any
}

type Engine struct {
	ServerOption []grpc.ServerOption
	Server       *grpc.Server
	ServiceList  []RegisterServiceOption
}

func (e *Engine) AddServerOption(opt ...grpc.ServerOption) {
	e.ServerOption = append(e.ServerOption, opt...)
}

func (e *Engine) RegisterService(serviceDesc *grpc.ServiceDesc, serviceServer any) {
	e.ServiceList = append(e.ServiceList, RegisterServiceOption{
		ServiceDesc:   serviceDesc,
		ServiceServer: serviceServer,
	})
}

func (e *Engine) Run(addr string) error {
	if len(addr) < 1 {
		panic("Please input an address to start")
	}
	if e.Server == nil {
		e.Server = grpc.NewServer(e.ServerOption...)
	}
	if len(e.ServiceList) > 0 {
		for _, s := range e.ServiceList {
			if t, ok := s.ServiceServer.(interface{ testEmbeddedByValue() }); ok {
				t.testEmbeddedByValue()
			}
			e.Server.RegisterService(s.ServiceDesc, s.ServiceServer)
		}
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		lis.Close()
		return err
	}
	defer func() {
		e.Server.Stop()
		lis.Close()
	}()
	err = e.Server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func New() *Engine {
	e := Engine{
		ServerOption: []grpc.ServerOption{
			grpc.UnaryInterceptor(LoggerInterceptor()),
		},
	}
	return &e
}
