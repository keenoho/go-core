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

type ControllerHandler func(ctx *gin.Context)

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
	App        *App
	PrefixPath string
	RouteMap   map[string]ControllerRouteMapStruct
}

func (c *Controller) URLMapping() {
	// empty, just for interface
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
	c.RegisterHttpHandler(c.App)
}

func (c *Controller) RegisterHttpHandler(app *App) {
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
		case "ANY":
			ginRouter.Any(value.Path, c.ControllerToGinHandler(ControllerHandler(value.Handler)))
		}
	}
}

func (c *Controller) ControllerToGinHandler(controller ControllerHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		controller(ctx)
	}
}

/**
 * @Params: data any, code int64, msg string, status int
 **/
func (c *Controller) SendJson(ctx *gin.Context, args ...any) {
	resData, status := MakeResponse(args...)
	if c.App.Type == APP_TYPE_HTTP {
		ctx.Header("Cache-Control", "no-cache")
		ctx.JSON(status, resData)
	}
}

/**
 * @Params: data string, status int
 **/
func (c *Controller) SendText(ctx *gin.Context, args ...any) {
	resData, status := MakeResponse(args...)
	if c.App.Type == APP_TYPE_HTTP {
		ctx.Header("Cache-Control", "no-cache")
		ctx.String(status, resData.Data.(string))
	}
}

func (c *Controller) BindParams(ctx *gin.Context, obj any) {
	if obj == nil {
		return
	}

	err := ctx.Copy().ShouldBind(obj)
	if err != nil {
		panic(ErrorData{Code: CODE_PARAMS_MISSING})
	}
}
