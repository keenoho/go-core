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
	Id         string
	Name       string
	Addr       string
	Server     *MicroServiceServer
	GrpcServer *grpc.Server
	RouteMap   map[string]MicroServiceControllerFunc
}

func (ms *MicroService) RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc) {
	if ms.RouteMap == nil {
		ms.RouteMap = make(map[string]MicroServiceControllerFunc)
	}
	ms.RouteMap[key] = handler
}

func (ms *MicroService) Run(addr string) error {
	ms.Addr = addr
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	ms.GrpcServer = grpc.NewServer(
	// grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			err = r
	// 		}
	// 	}()
	// 	return handler(ctx, req)
	// }),
	)
	ms.Server = &MicroServiceServer{
		Service: ms,
	}
	RegisterServiceHandlerServer(ms.GrpcServer, ms.Server)

	defer ms.GrpcServer.Stop()
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

	return nil
}
