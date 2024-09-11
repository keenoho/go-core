package core_test

import (
	"context"
	"fmt"

	"github.com/keenoho/go-core"
	"github.com/keenoho/go-core/protobuf"
)

// grpc
type TestGrpcController struct {
	core.GrpcController
	protobuf.UnimplementedBaseServiceServer
}

func (c *TestGrpcController) ServiceMapping() {
	c.Mapping(&protobuf.BaseService_ServiceDesc)
}

func (c *TestGrpcController) BaseRequest(ctx context.Context, in *protobuf.BaseRequestBody) (resp *protobuf.BaseResponseBody, err error) {
	fmt.Println(in)
	return resp, err
}
