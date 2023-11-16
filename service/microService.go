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
	Server         *MicroServiceServer
	GrpcServer     *grpc.Server
	RouteMap       map[string]MicroServiceControllerFunc
	MiddlewareList []MicroServiceMiddleware
}

func (ms *MicroService) RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc) {
	if ms.RouteMap == nil {
		ms.RouteMap = make(map[string]MicroServiceControllerFunc)
	}
	ms.RouteMap[key] = handler
}

func (ms *MicroService) Use(middlewares ...MicroServiceMiddleware) {
	if ms.MiddlewareList == nil {
		ms.MiddlewareList = []MicroServiceMiddleware{}
	}
	if len(middlewares) > 0 {
		for i := 0; i < len(middlewares); i++ {
			ms.MiddlewareList = append(ms.MiddlewareList, middlewares[i])
		}
	}
}

func (ms *MicroService) Run(addr string) error {
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
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	defer ms.GrpcServer.Stop()

	return nil
}
