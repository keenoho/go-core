package core

import (
	grpc "google.golang.org/grpc"
	"net"
)

var MicroServiceDebugMode = "debug"
var MicroServiceReleaseMode = "release"
var MicroServiceMode = "release"

type MicroServiceInterface interface {
	Run(addr string)
}

type MicroService struct {
	Server     *MicroServiceServer
	GrpcServer *grpc.Server
	RouteMap   map[string]MicroServiceControllerFunc
}

func (ms *MicroService) SetCustomServer(server *MicroServiceServer) {
	ms.Server = server
}

func (ms *MicroService) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	if ms.GrpcServer == nil {
		ms.GrpcServer = grpc.NewServer()
	}
	if ms.Server == nil {
		server := MicroServiceServer{
			Service: ms,
		}
		ms.Server = &server
	}
	RegisterServiceMsgHandlerServer(ms.GrpcServer, ms.Server)
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	defer ms.GrpcServer.Stop()

	return nil
}

func (ms *MicroService) RegisterRouteControllerFunc(key string, handler MicroServiceControllerFunc) {
	if ms.RouteMap == nil {
		ms.RouteMap = make(map[string]MicroServiceControllerFunc)
	}
	ms.RouteMap[key] = handler
}

func NewMicroService() *MicroService {
	ms := MicroService{
		RouteMap: make(map[string]MicroServiceControllerFunc),
	}
	return &ms
}

func SetMicroServiceMode(mode string) {
	MicroServiceMode = mode
}

func CreateMicroApp() *MicroService {
	conf := GetConfig()
	if conf["Env"] == "production" {
		SetMicroServiceMode(MicroServiceReleaseMode)
	} else {
		SetMicroServiceMode(MicroServiceDebugMode)
	}
	app := NewMicroService()
	return app
}
