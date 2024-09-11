package core

import (
	"google.golang.org/grpc"
)

type GrpcControllerInterface interface {
	Init(app *App)
	ServiceMapping()
	Register(ic GrpcControllerInterface)
}

type GrpcControllerRouteMapStruct struct {
	ServiceDesc *grpc.ServiceDesc
}

type GrpcController struct {
	App      *App
	RouteMap map[string]GrpcControllerRouteMapStruct
}

func (c *GrpcController) ServiceMapping() {
	// empty, just for interface
	/**
	such as:
	c.Mapping(&protobuf.BaseService_ServiceDesc)
	*/
}

func (c *GrpcController) Init(app *App) {
	c.App = app
	c.RouteMap = make(map[string]GrpcControllerRouteMapStruct)
}

func (c *GrpcController) Mapping(desc *grpc.ServiceDesc) {
	key := desc.ServiceName
	value := GrpcControllerRouteMapStruct{ServiceDesc: desc}
	c.RouteMap[key] = value
}

func (c *GrpcController) Register(ic GrpcControllerInterface) {
	for _, v := range c.RouteMap {
		c.App.RegisterGrpcService(v.ServiceDesc, ic)
	}
}
