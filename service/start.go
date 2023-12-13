package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/keenoho/go-core"
	"google.golang.org/grpc"
)

func CreateServiceApp(gprcServerOption ...grpc.ServerOption) *MicroService {
	serviceId := core.GetOneConfig("ServiceId")
	serviceName := core.GetOneConfig("ServiceName")

	if len(serviceId) < 1 {
		i4 := rand.Intn(9999)
		umi := time.Now().UnixMilli()
		serviceId = fmt.Sprintf("%d-%d", i4, umi)
	}
	if len(serviceName) < 1 {
		serviceName = "unknow"
	}

	app := &MicroService{
		Id:               serviceId,
		Name:             serviceName,
		RouteMap:         make(map[string]MicroServiceControllerFunc),
		GrpcServerOption: gprcServerOption,
	}

	return app
}
