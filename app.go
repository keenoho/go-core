package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core/microService"
)

func CreateApp(middlewares ...gin.HandlerFunc) *gin.Engine {
	conf := GetConfig()
	if conf["Env"] == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	app := gin.New()
	if _, hasTrustedProxies := conf["TrustedProxies"]; hasTrustedProxies {
		app.SetTrustedProxies([]string{conf["TrustedProxies"]})
	} else {
		app.SetTrustedProxies([]string{"*"})
	}
	app.Static(conf["StaticPath"], conf["StaticDir"])
	app.Use(AppErrorMiddleware())
	app.Use(AppLoggerMiddleware())
	app.Use(AppCorsMiddleware())

	if len(middlewares) > 0 {
		for _, m := range middlewares {
			app.Use(m)
		}
	}

	app.NoRoute(AppNotFoundMiddleware())

	return app
}

func CreateMicroApp() *microService.MicroService {
	app := microService.New()
	return app
}

func GetAppStartUpAddress() string {
	conf := GetConfig()
	return fmt.Sprintf("%s:%s", conf["Host"], conf["Port"])
}
