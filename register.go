package core

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

var RegisterCenterDebugMode = "debug"
var RegisterCenterReleaseMode = "release"
var RegisterCenterMode = "release"
var RC *RegisterCenter

type RegisterService struct {
	Name          string
	Id            string
	Key           string
	Address       string
	Health        bool
	LastCheckTime time.Time
}

type RegisterCenter struct {
	Client     *clientv3.Client
	ServiceMap map[string][]*RegisterService
}

func (rc *RegisterCenter) Print(printType string, format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("[RegisterCenter-"+printType+"] "+format, values...)
}

func (rc *RegisterCenter) debugPrint(format string, values ...any) {
	if RegisterCenterMode == RegisterCenterDebugMode {
		rc.Print("debug", format, values...)
	}
}

func (rc *RegisterCenter) errorPrint(format string, values ...any) {
	rc.Print("error", format, values...)
}

func (rc *RegisterCenter) Init(option ...clientv3.Op) {
	addr := GetRegisterAddress()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{addr},
		DialKeepAliveTimeout: time.Minute * 5,
		DialTimeout:          time.Second * 10,
	})
	if err != nil {
		rc.errorPrint("create client error: %v", err)
		panic(err)
	}
	rc.Client = cli
	defer func() {
		cli.Close()
		rc.Client = nil
	}()

	rc.initAllExistService()
	rc.watchServiceKey()
}

func (rc *RegisterCenter) getServicePrefixKey() string {
	conf := GetConfig("RegisterServicePrefixKey")
	servicePrefixKey := "service/"
	if len(conf["RegisterServicePrefixKey"]) > 0 {
		servicePrefixKey = conf["RegisterServicePrefixKey"]
	}
	return servicePrefixKey
}

func (rc *RegisterCenter) initAllExistService() {
	servicePrefixKey := rc.getServicePrefixKey()
	resp, err := rc.Client.Get(context.Background(), servicePrefixKey, clientv3.WithPrefix())
	if err != nil {
		rc.errorPrint("init service error: %v", err)
		panic(err)
	}
	for _, kv := range resp.Kvs {
		rc.addServiceByKv(kv)
	}
}

func (rc *RegisterCenter) addServiceByKv(kv *mvccpb.KeyValue) {
	key := string(kv.Key)
	value := string(kv.Value)
	content := strings.Split(key, "/")

	if len(content) != 3 {
		rc.errorPrint("service %v is not a correct key", key)
		return
	}

	name := content[1]
	id := content[2]

	for _, rs := range rc.ServiceMap[name] {
		if rs.Key == key {
			return
		}
	}

	service := RegisterService{
		Name:    name,
		Id:      id,
		Key:     key,
		Address: value,
	}
	res, err := MicroServiceOnceRequest(value, &ServiceMsgRequest{
		Url: "/common/health",
	})
	if err != nil || res.Code != 0 {
		rc.errorPrint("service [%v] check health error: %v, %v", key, err, res)
		return
	}
	rc.debugPrint("check health [%v] is ok", key)
	service.Health = true
	service.LastCheckTime = time.Now()
	rc.ServiceMap[name] = append(rc.ServiceMap[name], &service)
}

func (rc *RegisterCenter) removeServiceByKv(kv *mvccpb.KeyValue) {
	key := string(kv.Key)
	content := strings.Split(key, "/")
	if len(content) != 3 {
		rc.errorPrint("service %v is not a correct key", key)
		return
	}

	name := content[1]

	for idx, rs := range rc.ServiceMap[name] {
		if rs.Key == key {
			rc.ServiceMap[name] = append(rc.ServiceMap[name][:idx], rc.ServiceMap[name][idx+1:]...)
			return
		}
	}
}

func (rc *RegisterCenter) watchServiceKey() {
	servicePrefixKey := rc.getServicePrefixKey()
	watchCh := rc.Client.Watch(context.Background(), servicePrefixKey, clientv3.WithPrefix())
	for resp := range watchCh {
		for _, ev := range resp.Events {
			if ev.Type.String() == "PUT" {
				rc.addServiceByKv(ev.Kv)
			} else if ev.Type.String() == "DELETE" {
				rc.removeServiceByKv(ev.Kv)
			}
		}
	}
}

func (rc *RegisterCenter) GetService(serviceName string) (*RegisterService, error) {
	rsArr, isExist := rc.ServiceMap[serviceName]
	if !isExist {
		return nil, fmt.Errorf("the service %s is not exist", serviceName)
	}
	if len(rsArr) < 1 {
		return nil, fmt.Errorf("the service %s has no one", serviceName)
	}
	var result *RegisterService
	for _, rs := range rsArr {
		if result == nil {
			result = rs
			continue
		}
		if rs.Health && !result.Health {
			result = rs
		}
		if rs.Health && rs.LastCheckTime.After(result.LastCheckTime) {
			result = rs
		}
	}
	return result, nil
}

func SetRegisterCenterMode(mode string) {
	RegisterCenterMode = mode
}

func LoadRegisterCenter() {
	conf := GetConfig()
	if conf["Env"] == "production" {
		SetRegisterCenterMode(RegisterCenterReleaseMode)
	} else {
		SetRegisterCenterMode(RegisterCenterDebugMode)
	}
	RC = &RegisterCenter{
		ServiceMap: make(map[string][]*RegisterService),
	}
	go RC.Init()
}

func RegisterGetService(serviceName string) (*RegisterService, error) {
	return RC.GetService(serviceName)
}
