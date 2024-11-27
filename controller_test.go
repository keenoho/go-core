package core_test

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keenoho/go-core"
	"github.com/keenoho/go-core/protobuf"
)

// http
type TestHttpController struct {
	core.Controller
}

func (c *TestHttpController) Configure() {
	c.URLMapping("/mytest", "GET", c.MyTest)
}

func (c *TestHttpController) MyTest(ctx *gin.Context) {
	c.SendJson(ctx)
}

// grpc
type TestGrpcController struct {
	core.Controller
	protobuf.UnimplementedRouterServer
}

func (c *TestGrpcController) Configure() {
	c.ServiceDescMapping(&protobuf.Router_ServiceDesc)
}

func (c *TestGrpcController) RouterRequest(ctx context.Context, in *protobuf.RouterRequestBody) (resp *protobuf.RouterResponseBody, err error) {
	fmt.Println(in)
	return resp, err
}
