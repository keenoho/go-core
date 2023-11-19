package service

import (
	"fmt"
	"time"

	"github.com/keenoho/go-core"
	"github.com/keenoho/go-tool"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func CreateServiceApp() *MicroService {
	serviceId := core.GetOneConfig("ServiceId")
	serviceName := core.GetOneConfig("ServiceName")
	registerHost := core.GetOneConfig("RegisterHost")
	registerPort := core.GetOneConfig("RegisterPort")

	if len(serviceId) < 1 {
		serviceId = tool.UnionId16String()
	}
	if len(serviceName) < 1 {
		serviceName = tool.FileGetDirName("")
	}

	app := &MicroService{
		Id:       serviceId,
		Name:     serviceName,
		Addr:     "",
		RouteMap: make(map[string]MicroServiceControllerFunc),
	}

	serviceRegister := &MicroServiceRegister{
		Service: app,
		EtcdConfig: clientv3.Config{
			Endpoints:   []string{fmt.Sprintf("%s:%s", registerHost, registerPort)},
			DialTimeout: time.Second * 5,
		},
	}

	app.ServiceRegister = serviceRegister

	return app
}
