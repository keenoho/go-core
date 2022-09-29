package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 创建webapp
func CreateApp() *gin.Engine {
	conf := GetConfig()
	if conf.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	app := gin.New()
	app.SetTrustedProxies([]string{"*"})
	app.Static(conf.StaticPath, conf.StaticDir)
	app.Use(ErrorHandler())
	app.Use(LoggerMiddleware())
	app.Use(CorsMiddleware())
	app.Use(SessionMiddleware())
	app.NoRoute(NotFoundHandler())

	return app
}

// 服务启动地址
func StartUpAddress() string {
	conf := GetConfig()
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
