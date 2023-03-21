package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 创建web app
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

// 创建微服务app
func CreateMicroServiceApp() *MicroServiceApp {
	app := MicroServiceApp{}
	return &app
}

// 服务启动地址
func StartUpAddress() string {
	conf := GetConfig()
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
