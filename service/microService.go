package service

import (
	"net"
	reflect "reflect"

	"github.com/keenoho/go-core"
	grpc "google.golang.org/grpc"
)

type MicroServiceInterface interface {
	core.Logger
	RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc)
	Run(addr string) error
}

type MicroService struct {
	core.Logger
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
	typ := reflect.TypeOf(handler)
	handlerName := typ.Name()
	pkgPath := typ.PkgPath()
	ms.PrintDebug("%s\t\t--> %v", key, pkgPath+"/"+handlerName)
}

func (ms *MicroService) Use(middlewares ...MicroServiceMiddleware) {
	if ms.MiddlewareList == nil {
		ms.MiddlewareList = []MicroServiceMiddleware{}
	}
	for _, middleware := range middlewares {
		ms.MiddlewareList = append(ms.MiddlewareList, middleware)
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
	RegisterServiceMsgHandlerServer(ms.GrpcServer, ms.Server)
	ms.PrintDebug("Listening and serving HTTP on %s", addr)
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	defer ms.GrpcServer.Stop()

	return nil
}
