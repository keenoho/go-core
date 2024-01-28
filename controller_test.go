package core_test

import (
	"testing"

	"gitee.com/keenoho/go-core"
)

type MyController struct {
	core.Controller
}

func (c *MyController) URLMapping() {
	c.Mapping("/hi", "GET", c.MyMethod)
}

func (c *MyController) MyMethod(ctx *core.Context) core.ControllerResponse {
	return c.MakeResponse("hi")
}

func TestController(t *testing.T) {
	c := new(MyController)
	t.Log(c)
}
