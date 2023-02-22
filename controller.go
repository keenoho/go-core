package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// web controller
type WebControllerInterface interface {
	Init()
	URLMapping()
	Resigter(app *gin.Engine)
}

type RouteMapStruct struct {
	path   string
	method string
	fn     ControllerFunc
}

type WebController struct {
	PrefixPath string
	RouteMap   map[string]RouteMapStruct
}

// 业务路由添加到mapping
func (c *WebController) URLMapping() {
}

// 初始化
func (c *WebController) Init() {
	c.RouteMap = make(map[string]RouteMapStruct)
}

// 把路由加入到mapping
func (c *WebController) Mapping(path string, method string, fn ControllerFunc) {
	key := method + path
	value := RouteMapStruct{path: path, method: method, fn: fn}
	c.RouteMap[key] = value
}

// Controller转ginHandler
func ControllerToHandler(controller ControllerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, status := controller(ctx)
		ctx.JSON(status, data)
	}
}

// 注册controller到gin
func (c WebController) Resigter(app *gin.Engine) {
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
			ginRouter.GET(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		case "POST":
			ginRouter.POST(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		case "HEAD":
			ginRouter.HEAD(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		case "OPTIONS":
			ginRouter.OPTIONS(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		case "PUT":
			ginRouter.PUT(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		case "DELETE":
			ginRouter.DELETE(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		default:
			ginRouter.GET(value.path, ControllerToHandler(ControllerFunc(value.fn)))
		}

	}
}

// 注册controller
func ResigterController(app *gin.Engine, execControllers ...WebControllerInterface) {
	for _, execController := range execControllers {
		execController.Init()
		execController.URLMapping()
		execController.Resigter(app)
	}
}
