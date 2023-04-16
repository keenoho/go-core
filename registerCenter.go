package core

import (
	"context"
	"fmt"
	"github.com/keenoho/go-tool"
	"github.com/keenoho/go-tool/crypto"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var RegisterCenterDebugMode = "debug"
var RegisterCenterReleaseMode = "release"
var RegisterCenterMode = "release"
var RegisterCenterServicePrefixKey = "service"
var RegisterCenterCheckHealthTime = 60 * time.Second
var RC *RegisterCenter

type RegisterCenterOption struct {
	NotLoadCurrentService bool
	CurrentServiceAddress string
}

type RegisterCenter struct {
	Logger
	Client              *clientv3.Client
	RegisterServiceList []*RegisterService
	option              RegisterCenterOption
}

func (rc *RegisterCenter) Init() {
	addr := GetRegisterAddress()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{addr},
		DialKeepAliveTimeout: time.Minute * 5,
		DialTimeout:          time.Second * 10,
	})
	if err != nil {
		panic(err)
	}
	rc.Client = cli
	defer func() {
		cli.Close()
		rc.Client = nil
		rc.RegisterServiceList = make([]*RegisterService, 0)
	}()

	// init all existing service from store
	rc.initExistingService()

	// init current service
	rc.loadCurrentService()

	// watch service change from store
	rc.watchServiceChange()
}

func (rc *RegisterCenter) createServiceByKv(kv *mvccpb.KeyValue) (*RegisterService, error) {
	key := string(kv.Key)
	value := string(kv.Value)
	content := strings.Split(key, "/")

	if len(content) != 3 {
		return nil, fmt.Errorf("the key is not a correct format: %s", key)
	}

	app := content[1]
	id := content[2]

	rs := RegisterService{
		Key:     key,
		Value:   value,
		Address: value,
		App:     app,
		Id:      id,
	}
	rs.SetLoggerEnv(RegisterCenterMode)
	rs.SetLoggerName("RegisterService")
	return &rs, nil
}

func (rc *RegisterCenter) checkServiceIsInList(key string) int {
	for idx, rs := range rc.RegisterServiceList {
		if rs.Key == key {
			return idx
		}
	}
	return -1
}

func (rc *RegisterCenter) checkServiceIsInStore(key string, option ...clientv3.OpOption) *mvccpb.KeyValue {
	kvs, err := rc.GetKey(key, option...)
	if err != nil || len(kvs) < 1 {
		return nil
	}
	return kvs[0]
}

func (rc *RegisterCenter) putServiceByKv(kv *mvccpb.KeyValue, saveToStore bool) {
	key := string(kv.Key)
	value := string(kv.Value)
	rs, err := rc.createServiceByKv(kv)
	if err != nil {
		return
	}
	if rc.checkServiceIsInList(rs.Key) > -1 {
		return
	}
	rc.RegisterServiceList = append(rc.RegisterServiceList, rs)
	if saveToStore {
		rc.PutKey(key, value)
	}
}

func (rc *RegisterCenter) deleteServiceByKv(kv *mvccpb.KeyValue, saveToStore bool) {
	key := string(kv.Key)
	idx := rc.checkServiceIsInList(key)
	if idx > -1 {
		rc.RegisterServiceList = append(rc.RegisterServiceList[:idx], rc.RegisterServiceList[idx+1:]...)
	}
	if saveToStore {
		rc.DeleteKey(key)
	}
}

func (rc *RegisterCenter) initExistingService() {
	kvs, err := rc.GetKey(RegisterCenterServicePrefixKey+"/", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	for _, kv := range kvs {
		rc.putServiceByKv(kv, false)
	}
}

func (rc *RegisterCenter) loadCurrentService() {
	conf := GetConfig()
	if rc.option.NotLoadCurrentService {
		return
	}
	app := conf["App"]
	address := rc.option.CurrentServiceAddress
	if len(address) < 1 {
		address = fmt.Sprintf("%s:%s", tool.ServerInternalIp(), conf["Port"])
	}
	isInList := false
	isInStore := false
	var storeService *RegisterService
	var storeKv *mvccpb.KeyValue
	for _, rs := range rc.RegisterServiceList {
		if rs.App == app && rs.Address == address {
			isInList = true
			storeService = rs
			break
		}
	}
	kvs, _ := rc.GetKey(RegisterCenterServicePrefixKey+"/"+app, clientv3.WithPrefix())
	for _, kv := range kvs {
		v := string(kv.Value)
		if v == address {
			isInStore = true
			storeKv = kv
			break
		}
	}

	if isInList && isInStore {
		return
	} else if !isInList && isInStore {
		rc.putServiceByKv(storeKv, false)
	} else if isInList && !isInStore {
		idx := rc.checkServiceIsInList(storeService.Key)
		if idx > -1 {
			rc.RegisterServiceList = append(rc.RegisterServiceList[:idx], rc.RegisterServiceList[idx+1:]...)
		}
		storeKv.Key = []byte(storeService.Key)
		storeKv.Value = []byte(address)
		rc.putServiceByKv(storeKv, true)
	} else {
		id := strings.Split(crypto.UUID(), "-")[0]
		key := fmt.Sprintf("%s/%s/%s", RegisterCenterServicePrefixKey, app, id)
		storeKv = &mvccpb.KeyValue{
			Key:   []byte(key),
			Value: []byte(address),
		}
		rc.putServiceByKv(storeKv, true)
	}
}

func (rc *RegisterCenter) watchServiceChange() {
	rc.WatchKey(RegisterCenterServicePrefixKey+"/", func(even *clientv3.Event) {
		if even.Type.String() == "PUT" {
			rc.putServiceByKv(even.Kv, false)
		} else if even.Type.String() == "DELETE" {
			rc.deleteServiceByKv(even.Kv, false)
		}
	}, clientv3.WithPrefix())
}

func (rc *RegisterCenter) WatchKey(key string, callback func(even *clientv3.Event), option ...clientv3.OpOption) {
	watchCh := rc.Client.Watch(context.Background(), key, option...)
	for resp := range watchCh {
		for _, ev := range resp.Events {
			callback(ev)
		}
	}
}

func (rc *RegisterCenter) GetKey(key string, option ...clientv3.OpOption) ([]*mvccpb.KeyValue, error) {
	resp, err := rc.Client.Get(context.Background(), key, option...)
	if err != nil {
		return nil, err
	}
	return resp.Kvs, nil
}

func (rc *RegisterCenter) PutKey(key string, value string, option ...clientv3.OpOption) error {
	_, err := rc.Client.Put(context.Background(), key, value, option...)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegisterCenter) DeleteKey(key string, option ...clientv3.OpOption) error {
	_, err := rc.Client.Delete(context.Background(), key, option...)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegisterCenter) GetService(app string) (*RegisterService, error) {
	rsList := make([]*RegisterService, 0)
	for _, rs := range rc.RegisterServiceList {
		if rs.App == app {
			rsList = append(rsList, rs)
		}
	}
	if len(rsList) > 0 {
		return rsList[0], nil
	}
	key := fmt.Sprintf("%s/%s", RegisterCenterServicePrefixKey, app)
	kvs, err := rc.GetKey(key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if len(kvs) < 1 {
		return nil, fmt.Errorf("no service exist: %s", app)
	}
	rs, err := rc.createServiceByKv(kvs[0])
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (rc *RegisterCenter) RequestService(app string, requestData *ServiceMsgRequest) (*ServiceMsgResponse, error) {
	rs, err := rc.GetService(app)
	if err != nil {
		return nil, err
	}
	if !rs.Health {
		res := rs.CheckHealth()
		if !res {
			return nil, fmt.Errorf("service is not health: %s", app)
		}
	}
	return MicroServiceClientOnceRequest(rs.Address, requestData)
}

func SetRegisterCenterMode(mode string) {
	RegisterCenterMode = mode
}

func LoadRegisterCenter(loadOption ...RegisterCenterOption) {
	conf := GetConfig()
	if conf["Env"] == "production" {
		SetRegisterCenterMode(RegisterCenterReleaseMode)
	} else {
		SetRegisterCenterMode(RegisterCenterDebugMode)
	}
	if len(conf["RegisterServicePrefixKey"]) > 0 {
		RegisterCenterServicePrefixKey = conf["RegisterServicePrefixKey"]
	}
	option := RegisterCenterOption{}
	if len(loadOption) > 0 {
		for _, o := range loadOption {
			option.NotLoadCurrentService = o.NotLoadCurrentService
			option.CurrentServiceAddress = o.CurrentServiceAddress
		}
	}
	RC = &RegisterCenter{
		option: option,
	}
	RC.SetLoggerEnv(RegisterCenterMode)
	RC.SetLoggerName("RegisterCenter")
	RC.Init()
}

func RequestRegisterService(app string, requestData *ServiceMsgRequest) (*ServiceMsgResponse, error) {
	if RC == nil || RC.Client == nil {
		return nil, fmt.Errorf("register center is not init")
	}
	return RC.RequestService(app, requestData)
}
