package service

import (
	"github.com/keenoho/go-core"
	"github.com/keenoho/go-tool"
)

func CreateServiceApp() *MicroService {
	serviceId := core.GetOneConfig("ServiceId")
	serviceName := core.GetOneConfig("ServiceName")

	if len(serviceId) < 1 {
		serviceId = tool.UnionId16String()
	}
	if len(serviceName) < 1 {
		serviceName = tool.FileGetDirName("")
	}

	app := &MicroService{
		Id:       serviceId,
		Name:     serviceName,
		RouteMap: make(map[string]MicroServiceControllerFunc),
	}

	return app
}
