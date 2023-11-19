package service

import (
	"net"

	grpc "google.golang.org/grpc"
)

type MicroServiceInterface interface {
	RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc)
	Run(addr string) error
}

type MicroService struct {
	Id              string
	Name            string
	Addr            string
	Server          *MicroServiceServer
	GrpcServer      *grpc.Server
	RouteMap        map[string]MicroServiceControllerFunc
	ServiceRegister *MicroServiceRegister
}

func (ms *MicroService) RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc) {
	if ms.RouteMap == nil {
		ms.RouteMap = make(map[string]MicroServiceControllerFunc)
	}
	ms.RouteMap[key] = handler
}

func (ms *MicroService) InitServiceRegister() {
	if ms.ServiceRegister != nil {
		go ms.ServiceRegister.Init(ms.Name, ms.Id, ms.Addr)
	}
}

func (ms *MicroService) Run(addr string) error {
	ms.Addr = addr
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	ms.GrpcServer = grpc.NewServer()
	ms.Server = &MicroServiceServer{
		Service: ms,
	}
	RegisterServiceHandlerServer(ms.GrpcServer, ms.Server)

	ms.InitServiceRegister()

	defer ms.GrpcServer.Stop()
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
