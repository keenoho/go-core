package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type AppControllerFunc func(ctx *gin.Context) (ResponseData, int)

type AppControllerInterface interface {
	Init()
	URLMapping()
	Register(app *gin.Engine)
}

type AppRouteMapStruct struct {
	path   string
	method string
	fn     AppControllerFunc
}

type AppController struct {
	PrefixPath string
	RouteMap   map[string]AppRouteMapStruct
}

func (c *AppController) URLMapping() {
}

func (c *AppController) Init() {
	c.RouteMap = make(map[string]AppRouteMapStruct)
}

func (c *AppController) Mapping(path string, method string, fn AppControllerFunc) {
	key := method + ":" + path
	value := AppRouteMapStruct{path: path, method: method, fn: fn}
	c.RouteMap[key] = value
}

func (c *AppController) Register(app *gin.Engine) {
	var ginRouter *gin.RouterGroup
	if len(c.PrefixPath) > 0 {
		ginRouter = app.Group(c.PrefixPath)
	} else {
		ginRouter = app.Group("")
	}
	for key := range c.RouteMap {
		value := c.RouteMap[key]
		switch strings.ToUpper(value.method) {
		case "GET":
			ginRouter.GET(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		case "POST":
			ginRouter.POST(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		case "HEAD":
			ginRouter.HEAD(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		case "OPTIONS":
			ginRouter.OPTIONS(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		case "PUT":
			ginRouter.PUT(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		case "DELETE":
			ginRouter.DELETE(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		default:
			ginRouter.GET(value.path, webControllerToHandler(AppControllerFunc(value.fn)))
		}

	}
}

func (c *AppController) BindParams(ctx *gin.Context, queryBind any) {
	if queryBind != nil {
		err := ctx.Copy().ShouldBind(queryBind)
		if err != nil {
			panic(ErrorData{Code: CODE_PARAMS_MISSING})
		}
	}
}

func webControllerToHandler(controller AppControllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := controller(ctx)
		ctx.Header("Cache-Control", "no-cache")
		ctx.JSON(status, data)
	}
}

func RegisterAppController(app *gin.Engine, execControllers ...AppControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Register(app)
	}
}
