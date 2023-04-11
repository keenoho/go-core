package core

import (
	context "context"
)

type MicroServiceControllerFunc func(ctx context.Context, in *ServiceMsgRequest) ServiceResponseData

type MicroServiceControllerInterface interface {
	Init()
	URLMapping()
	Register(app *MicroService)
}

type MicroServiceRouteMapStruct struct {
	path string
	fn   MicroServiceControllerFunc
}

type MicroServiceController struct {
	RouteMap map[string]MicroServiceRouteMapStruct
}

func (c *MicroServiceController) URLMapping() {
}

func (c *MicroServiceController) Init() {
	c.RouteMap = make(map[string]MicroServiceRouteMapStruct)
}

func (c *MicroServiceController) Mapping(path string, fn MicroServiceControllerFunc) {
	value := MicroServiceRouteMapStruct{path: path, fn: fn}
	c.RouteMap[path] = value
}

func (c *MicroServiceController) Register(app *MicroService) {
	for _, handler := range c.RouteMap {
		app.RegisterRouteControllerFunc(handler.path, handler.fn)
	}
}

func RegisterMicroServiceController(app *MicroService, execControllers ...MicroServiceControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Register(app)
	}
}
