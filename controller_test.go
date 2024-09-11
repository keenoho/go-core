package core_test

import (
	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
)

// http
type TestHttpController struct {
	core.Controller
}

func (c *TestHttpController) URLMapping() {
	c.Mapping("/mytest", "GET", c.MyTest)
}

func (c *TestHttpController) MyTest(ctx *gin.Context) {
	c.SendJson(ctx)
}
