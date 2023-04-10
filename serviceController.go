package core

import (
	context "context"
)

type ServiceControllerFunc func(ctx context.Context, in *ServiceMsgRequest) ServiceResponseData

type ServiceControllerInterface interface {
	Init()
	URLMapping()
	Register(app *MicroService)
}

type ServiceRouteMapStruct struct {
	path string
	fn   ServiceControllerFunc
}

type ServiceController struct {
	RouteMap map[string]ServiceRouteMapStruct
}

func (c *ServiceController) URLMapping() {
}

func (c *ServiceController) Init() {
	c.RouteMap = make(map[string]ServiceRouteMapStruct)
}

func (c *ServiceController) Mapping(path string, fn ServiceControllerFunc) {
	value := ServiceRouteMapStruct{path: path, fn: fn}
	c.RouteMap[path] = value
}

func (c *ServiceController) Register(app *MicroService) {
	for _, handler := range c.RouteMap {
		app.RegisterRouteControllerFunc(handler.path, handler.fn)
	}
}

func (c *ServiceController) MakeServiceResponse(args ...any) ServiceResponseData {
	resData := ServiceResponseData{
		Data: nil,
		Code: 0,
		Msg:  "ok",
	}
	for i, v := range args {
		switch i {
		case 0:
			{
				resData.Data = v
				break
			}
		case 1:
			{
				resData.Code = int64(v.(int))
				break
			}
		case 2:
			{
				resData.Msg = v.(string)
				break
			}
		}
	}
	return resData
}

func RegisterServiceController(app *MicroService, execControllers ...ServiceControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Register(app)
	}
}
