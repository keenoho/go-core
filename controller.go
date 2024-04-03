package core

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type ControllerResponse struct {
	Body   ResponseData
	Status int
}

type ControllerHandler func(ctx *Context) ControllerResponse

type ControllerInterface interface {
	Init(app *App)
	URLMapping()
	Register()
}

type ControllerRouteMapStruct struct {
	Path    string
	Method  string
	Handler ControllerHandler
}

type Controller struct {
	Module
	App        *App
	PrefixPath string
	RouteMap   map[string]ControllerRouteMapStruct
}

func (c *Controller) URLMapping() {
}

func (c *Controller) Init(app *App) {
	c.App = app
	c.RouteMap = make(map[string]ControllerRouteMapStruct)
}

func (c *Controller) SetPrefixPath(prefixPath string) {
	c.PrefixPath = prefixPath
}

func (c *Controller) Mapping(path string, method string, handler ControllerHandler) {
	key := fmt.Sprintf("%s:%s", method, path)
	value := ControllerRouteMapStruct{Path: path, Method: method, Handler: handler}
	c.RouteMap[key] = value
}

func (c *Controller) Register() {
	switch c.App.Type {
	case APP_TYPE_HTTP:
		{
			c.registerHttpHandler(c.App)
			break
		}
	case APP_TYPE_GRPC:
		{
			c.RegisterGrpcHandler(c.App)
			break
		}
	}
}

func (c *Controller) registerHttpHandler(app *App) {
	var ginRouter *gin.RouterGroup
	if len(c.PrefixPath) > 0 {
		ginRouter = app.HttpServer.Group(c.PrefixPath)
	} else {
		ginRouter = app.HttpServer.Group("")
	}
	for key := range c.RouteMap {
		value := c.RouteMap[key]
		switch strings.ToUpper(value.Method) {
		case "GET":
			ginRouter.GET(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		case "POST":
			ginRouter.POST(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		case "HEAD":
			ginRouter.HEAD(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		case "OPTIONS":
			ginRouter.OPTIONS(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		case "PUT":
			ginRouter.PUT(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		case "DELETE":
			ginRouter.DELETE(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		default:
			ginRouter.GET(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		}
	}
}

func (c *Controller) RegisterGrpcHandler(app *App) {

}

func (c *Controller) ControllerToGinHandler(controller ControllerHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := controller(&Context{App: c.App, HttpServiceContext: ctx})
		ctx.Header("Cache-Control", "no-cache")
		ctx.JSON(res.Status, res.Body)
	}
}

func (c *Controller) ControllerToGrpcServiceHandler(controller ControllerHandler) {

}

/**
 * @Params: data any, code int64, msg string, status int
 * @Response: ControllerResponse
 **/
func (c *Controller) MakeResponse(args ...any) ControllerResponse {
	resData, status := MakeResponse(args...)
	return ControllerResponse{
		Body:   resData,
		Status: status,
	}
}

func (c *Controller) BindParams(ctx *Context, obj any) {
	if obj == nil {
		return
	}

	switch c.App.Type {
	case APP_TYPE_HTTP:
		{
			err := ctx.HttpServiceContext.Copy().ShouldBind(obj)
			if err != nil {
				tags := FilterOutMissingTags(err.Error())
				code := CODE_PARAMS_MISSING
				msg := CODE_MESSAGE_MAP[code]
				if len(tags) > 0 {
					msg += ":" + strings.Join(tags, ",")
				}
				panic(ErrorData{Code: code, Msg: msg})
			}
			break
		}
	case APP_TYPE_GRPC:
		{
			break
		}
	}
}

func RegisterController(app *App, execController any) {
	execController.(ControllerInterface).Init(app)
	execController.(ControllerInterface).URLMapping()
	execController.(ControllerInterface).Register()
}
