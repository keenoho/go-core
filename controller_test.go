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
	protobuf.UnimplementedBaseServer
}

func (c *TestGrpcController) Configure() {
	c.ServiceDescMapping(&protobuf.Base_ServiceDesc)
}

func (c *TestGrpcController) BaseRequest(ctx context.Context, in *protobuf.BaseRequestBody) (resp *protobuf.BaseResponseBody, err error) {
	fmt.Println(in)
	return resp, err
}
