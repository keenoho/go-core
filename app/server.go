package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core/common"
	"github.com/keenoho/go-core/middleware"
)

func CreateServer(middlewares ...gin.HandlerFunc) *gin.Engine {
	conf := common.GetConfig()
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
	app.Use(middleware.AppErrorMiddleware())
	app.Use(middleware.AppLoggerMiddleware())
	app.Use(middleware.AppCorsMiddleware())

	if len(middlewares) > 0 {
		for _, m := range middlewares {
			app.Use(m)
		}
	}

	app.NoRoute(middleware.AppNotFoundMiddleware())

	return app
}

func GetServerStartUpAddress() string {
	conf := common.GetConfig()
	return fmt.Sprintf("%s:%s", conf["Host"], conf["Port"])
}
