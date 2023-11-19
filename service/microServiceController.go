package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/keenoho/go-core"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (c *MicroServiceController) SendServiceRequest(ctx *MicroServiceContext, target string, url string, data any) (*ServiceResponse, error) {
	jsonBytes, _ := json.Marshal(data)
	requestData := ServiceRequest{
		Url:  url,
		Data: jsonBytes,
	}
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := NewServiceHandlerClient(conn)
	grpcCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Send(grpcCtx, &requestData)
	return res, err
}

func RegisterMicroServiceController(ms *MicroService, execControllers ...MicroServiceControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Register(ms)
	}
}
