package service

import (
	"encoding/json"

	"github.com/keenoho/go-core"
)

type MicroServiceControllerFunc func(ctx *MicroServiceContext) ResponseData

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
	Prefix   string
	RouteMap map[string]MicroServiceRouteMapStruct
}

func (c *MicroServiceController) URLMapping() {
}

func (c *MicroServiceController) Init() {
	c.RouteMap = make(map[string]MicroServiceRouteMapStruct)
}

func (c *MicroServiceController) Mapping(path string, fn MicroServiceControllerFunc) {
	route := MicroServiceRouteMapStruct{path: path, fn: fn}
	c.RouteMap[path] = route
}

func (c *MicroServiceController) Register(ms *MicroService) {
	for _, route := range c.RouteMap {
		key := c.Prefix + route.path
		ms.RegisterRouteControllerFunc(key, route.fn)
	}
}

func (c *MicroServiceController) BindJson(ctx *MicroServiceContext, paramsBind any) {
	err := json.Unmarshal(ctx.RequestIn.Data, &paramsBind)
	if err != nil {
		panic(core.ErrorData{Code: core.CODE_PARAMS_MISSING})
	}
}

func RegisterMicroServiceController(ms *MicroService, execControllers ...MicroServiceControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Register(ms)
	}
}
