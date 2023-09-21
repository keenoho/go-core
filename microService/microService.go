package service

import (
	grpc "google.golang.org/grpc"
	"net"
	"reflect"
)

var MicroServiceDebugMode = "debug"
var MicroServiceReleaseMode = "release"
var MicroServiceMode = "release"

type MicroServiceInterface interface {
	Run(addr string)
}

type MicroService struct {
	Logger
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
		server.SetLoggerEnv(MicroServiceMode)
		server.SetLoggerName("MicroServiceServer")
		ms.Server = &server
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
	app.SetLoggerEnv(MicroServiceMode)
	app.SetLoggerName("MicroService")
	return app
}
