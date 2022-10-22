# go-core

A web framework for our team and project, framework base on gin.

## Install:

```shell
go get -u github.com/keenoho/go-core
```

## Usage

1. create a controller file
``` go
// demo.controller.go

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
)

type DemoController struct {
	core.WebController
}

func (c *DemoController) URLMapping() {
	c.Mapping("/somePath", "POST", c.Demo)
}

func (c *DemoController) Demo(ctx *gin.Context) (core.ResponseData, int) {
    responseData := "hello"
    responseCode := 0
    responseMsg := "msg"
	return core.MakeResponse(responseData, responseCode, responseMsg)
}
```
2. register the controller
``` go
// router.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
)

func ResigterAllController(app *gin.Engine) {
	core.ResigterController(app, new(DemoController))
}
```
3. create an app and start it up
``` go
// main.go
package main

import (
	"github.com/keenoho/go-core"
)

func main() {
	core.LoadConfig()
	core.LoadDb()
	core.LoadRedis()

	serverStartUpAddress := core.StartUpAddress()

	app := core.CreateApp()
	ResigterAllController(app)
	app.Run(serverStartUpAddress)
}
```