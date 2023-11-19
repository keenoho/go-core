package service

import (
	"context"
	"fmt"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type MicroServiceRegistItem struct {
	Key         string
	ServiceName string
	ServiceId   string
	ServiceAddr string
}

type MicroServiceRegisterInterface interface {
}

type MicroServiceRegister struct {
	Service         *MicroService
	EtcdConfig      clientv3.Config
	EtcdClient      *clientv3.Client
	ServiceGroupMap map[string][]MicroServiceRegistItem
}

func (msr *MicroServiceRegister) etcdRespToServiceItem(resp *clientv3.GetResponse) []MicroServiceRegistItem {
	var result []MicroServiceRegistItem
	if resp.Count < 1 {
		return result
	}
	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		keySplits := strings.Split(key, ":")
		if len(keySplits) < 3 {
			continue
		}
		tmpItem := MicroServiceRegistItem{
			Key:         key,
			ServiceName: keySplits[1],
			ServiceId:   keySplits[2],
			ServiceAddr: string(kv.Value),
		}
		result = append(result, tmpItem)
	}
	return result
}

func (msr *MicroServiceRegister) serviceItemToGroupMap() {
	msr.ServiceGroupMap = map[string][]MicroServiceRegistItem{}
	allServiceItems := msr.GetEtcdService()

	for _, item := range allServiceItems {
		if len(item.ServiceName) < 1 {
			continue
		}
		_, hasKey := msr.ServiceGroupMap[item.ServiceName]
		if !hasKey {
			msr.ServiceGroupMap[item.ServiceName] = []MicroServiceRegistItem{item}
		} else {
			msr.ServiceGroupMap[item.ServiceName] = append(msr.ServiceGroupMap[item.ServiceName], item)
		}
	}
}

func (msr *MicroServiceRegister) PutEtcdService(service MicroServiceRegistItem) error {
	if len(service.Key) < 1 && len(service.ServiceName) > 0 && len(service.ServiceId) > 0 {
		service.Key = fmt.Sprintf(ETCD_SERVICE_REGIST_KEY, service.ServiceName, service.ServiceId)
	}

	if len(service.Key) < 1 {
		return fmt.Errorf("the key params is empty")
	}

	if len(service.ServiceAddr) < 1 {
		return fmt.Errorf("the addr params is empty")
	}

	_, err := msr.EtcdClient.Put(context.Background(), service.Key, service.ServiceAddr)
	return err
}

func (msr *MicroServiceRegister) DeleteEtcdService(service MicroServiceRegistItem) error {
	if len(service.Key) < 1 && len(service.ServiceName) > 0 && len(service.ServiceId) > 0 {
		service.Key = fmt.Sprintf(ETCD_SERVICE_REGIST_KEY, service.ServiceName, service.ServiceId)
	}

	if len(service.Key) < 1 {
		return fmt.Errorf("the key params is empty")
	}

	if len(service.ServiceAddr) < 1 {
		return fmt.Errorf("the addr params is empty")
	}

	_, err := msr.EtcdClient.Delete(context.Background(), service.Key)
	return err
}

func (msr *MicroServiceRegister) GetEtcdService(name ...string) []MicroServiceRegistItem {
	var serviceItem []MicroServiceRegistItem

	if len(name) > 0 {
		for _, n := range name {
			resp, err := msr.EtcdClient.Get(context.Background(), ETCD_SERVICE_REGIST_KEY_PREFIX+n, clientv3.WithPrefix())
			if err != nil {
				continue
			}
			serviceItem = append(serviceItem, msr.etcdRespToServiceItem(resp)...)
		}
	} else {
		resp, err := msr.EtcdClient.Get(context.Background(), ETCD_SERVICE_REGIST_KEY_PREFIX, clientv3.WithPrefix())
		if err != nil {
			return serviceItem
		}
		serviceItem = msr.etcdRespToServiceItem(resp)
	}

	return serviceItem
}

func (msr *MicroServiceRegister) GetService(name string) (MicroServiceRegistItem, error) {
	var result MicroServiceRegistItem
	serviceList, hasServieList := msr.ServiceGroupMap[name]
	if !hasServieList {
		return result, fmt.Errorf("the service name of %s is not register", name)
	}
	if len(serviceList) < 1 {
		return result, fmt.Errorf("the service name of %s has no one exist", name)
	}
	return serviceList[0], nil
}

func (msr *MicroServiceRegister) Init(name string, id string, addr string) {
	if msr.EtcdClient != nil {
		msr.EtcdClient.Close()
		msr.EtcdClient = nil
	}
	cli, err := clientv3.New(msr.EtcdConfig)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	msr.EtcdClient = cli

	msr.PutEtcdService(MicroServiceRegistItem{
		Key:         fmt.Sprintf(ETCD_SERVICE_REGIST_KEY, name, id),
		ServiceName: name,
		ServiceId:   id,
		ServiceAddr: addr,
	})

	msr.serviceItemToGroupMap()

	watchCh := cli.Watch(context.Background(), ETCD_SERVICE_REGIST_KEY_PREFIX, clientv3.WithPrefix())
	for wresp := range watchCh {
		for _, ev := range wresp.Events {
			if ev.Type == clientv3.EventTypePut || ev.Type == clientv3.EventTypeDelete {
				msr.serviceItemToGroupMap()
			}
		}
	}
}
