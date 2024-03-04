package core

import (
	"fmt"
	"reflect"

	"github.com/keenoho/go-core/grpc_engine"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

const APP_TYPE_HTTP = "http"
const APP_TYPE_GRPC = "grpc"

type AppOption struct {
	Id   string
	Type string
	Host string
	Port string
}

type App struct {
	Id         string
	Type       string
	Host       string
	Port       string
	HttpServer *gin.Engine
	GrpcServer *grpc_engine.Engine
}

func (app *App) Init() {
	// set env
	env := ConfigGet(FIELD_ENV)
	if env == "production" {
		SetMode(ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
		grpc_engine.SetMode(grpc_engine.ReleaseMode)
	} else {
		SetMode(DebugMode)
		gin.SetMode(gin.DebugMode)
		grpc_engine.SetMode(grpc_engine.DebugMode)
	}

	switch app.Type {
	case APP_TYPE_HTTP:
		{
			app.initHttpServer()
			break
		}
	case APP_TYPE_GRPC:
		{
			app.initGrpcServer()
			break
		}
	}
}

func (app *App) initHttpServer() {
	app.HttpServer = gin.New()
	trustedProxies := ConfigGet("TRUSTED_PROXIES")
	staticDir := ConfigGet("STATIC_DIR")
	staticPath := ConfigGet("STATIC_PATH")
	if len(trustedProxies) > 0 {
		app.HttpServer.SetTrustedProxies([]string{trustedProxies})
	} else {
		app.HttpServer.SetTrustedProxies([]string{"*"})
	}
	if len(staticDir) > 0 && len(staticPath) > 0 {
		app.HttpServer.Static(staticPath, staticDir)
	}
}

func (app *App) initGrpcServer() {
	app.GrpcServer = grpc_engine.New()
}

func (app *App) InitModule(modules ...ModuleInterface) {
	for _, module := range modules {
		module.Init(app)
	}
}

func (app *App) InitModuleUseCustomCaller(customCaller func(app *App, m any), modules ...ModuleInterface) {
	for _, module := range modules {
		customCaller(app, module)
	}
}

func (app *App) UseHttpMiddleware(middleware ...gin.HandlerFunc) {
	app.HttpServer.Use(middleware...)
}

func (app *App) AddGrpcServerOption(option ...grpc.ServerOption) {
	app.GrpcServer.AddServerOption(option...)
}

func (app *App) RegisterGrpcService(sd *grpc.ServiceDesc, ss any) {
	app.GrpcServer.RegisterService(sd, ss)
}

func (app *App) Start() error {
	startUpAddr := fmt.Sprintf("%s:%s", app.Host, app.Port)

	switch app.Type {
	case APP_TYPE_HTTP:
		{
			return app.HttpServer.Run(startUpAddr)
		}
	case APP_TYPE_GRPC:
		{
			return app.GrpcServer.Run(startUpAddr)
		}
	default:
		return nil
	}
}

func AppNew(options ...AppOption) *App {
	app := App{}
	newOptions := []AppOption{
		{
			Id:   DEFAULT_APP_ID,
			Type: DEFAULT_APP_TYPE,
			Host: DEFAULT_HOST,
			Port: DEFAULT_PORT,
		},
	}
	newOptions = append(newOptions, options...)

	for _, opt := range newOptions {
		StructConvert(opt, &app, func(fieldName string, fieldValue reflect.Value) bool {
			return len(fieldValue.String()) > 0
		})
	}
	app.Init()
	return &app
}
