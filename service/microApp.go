package service

func CreateMicroApp(middlewares ...MicroServiceMiddleware) *MicroService {
	app := &MicroService{
		RouteMap: make(map[string]MicroServiceControllerFunc),
	}
	app.SetLoggerName("MicroService")

	if len(middlewares) > 0 {
		for _, m := range middlewares {
			app.Use(m)
		}
	}

	return app
}
