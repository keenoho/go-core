package core_test

import (
	"testing"

	"github.com/keenoho/go-core"
)

func TestHttpServer(t *testing.T) {
	core.ConfigLoad()
	app := core.AppNew()
	app.RegisterController(
		new(TestHttpController),
	)
	err := app.Start()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGrpcServer(t *testing.T) {
	core.ConfigLoad()
	app := core.AppNew(core.AppOption{
		Type: core.APP_TYPE_GRPC,
		Port: "1234",
	})
	// app.RegisterGrpcService(&protobuf.BaseService_ServiceDesc, &TestGrpcController{})
	app.RegisterGrpcController(
		new(TestGrpcController),
	)
	err := app.Start()
	if err != nil {
		t.Fatal(err)
	}
}
