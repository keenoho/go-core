package extend

import "gitee.com/keenoho/go-core"

func UseDefaultHttpMiddleware(app *core.App) {
	app.HttpServer.Use(HttpErrorMiddleware(), HttpCorsMiddleware())
	app.HttpServer.NoRoute(HttpNotFoundMiddleware())
}
