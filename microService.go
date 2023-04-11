package core

import (
	"fmt"
	"net"
	"strings"

	grpc "google.golang.org/grpc"
)

type MicroServiceInterface interface {
	Run(addr string)
}

type MicroService struct {
	Server     *MicroServiceServer
	GrpcServer *grpc.Server
	RouteMap   map[string]MicroServiceControllerFunc
}

var DebugMode = "debug"
var ReleaseMode = "release"
var ServiceMode = "release"

func (ms *MicroService) Print(printType string, format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("[MicroService-"+printType+"] "+format, values...)
}

func (ms *MicroService) debugPrint(format string, values ...any) {
	if ServiceMode == DebugMode {
		ms.Print("debug", format, values...)
	}
}

func (ms *MicroService) errorPrint(format string, values ...any) {
	ms.Print("error", format, values...)
}

func (ms *MicroService) SetCustomServer(server *MicroServiceServer) {
	ms.Server = server
}

func (ms *MicroService) Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		ms.errorPrint("tcp listen error: %v", err)
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
	ms.debugPrint("Listening and serving on %s", addr)
	err = ms.GrpcServer.Serve(lis)
	if err != nil {
		ms.errorPrint("grpc serve error: %v", err)
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

func SetMode(mode string) {
	ServiceMode = mode
}

func CreateMicroApp() *MicroService {
	conf := GetConfig()
	if conf["Env"] == "production" {
		SetMode(ReleaseMode)
	} else {
		SetMode(DebugMode)
	}
	app := NewMicroService()
	return app
}
