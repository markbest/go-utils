package utils

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var EtcdClient *Etcd

type Etcd struct {
	client *clientv3.Client
}

//new conn
func NewEtcdConn(servers []string) *Etcd {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   servers,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	EtcdClient = &Etcd{client: cli}
	return EtcdClient
}

//create new node
func (e *Etcd) Put(key string, value string) error {
	_, err := e.client.Put(context.Background(), key, value)
	return err
}

//create new node with ttl
func (e *Etcd) PutWithTTL(key string, value string, leaseResp *clientv3.LeaseGrantResponse) error {
	_, err := e.client.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
	return err
}

//get node value
func (e *Etcd) Get(key string) (value string, err error) {
	resp, err := e.client.Get(context.Background(), key)
	if err != nil {
		if resp.Count > 0 {
			for _, k := range resp.Kvs {
				value = string(k.Value)
				break
			}
		}
	}
	return value, err
}

//get keys list by prefix
func (e *Etcd) GetKeysByPrefix(prefix string) (keys []string, err error) {
	resp, err := e.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return keys, err
	}
	if resp.Count > 0 {
		for _, k := range resp.Kvs {
			keys = append(keys, string(k.Key))
		}
	}
	return keys, err
}

//delete key
func (e *Etcd) Del(key string) (err error) {
	_, err = e.client.Delete(context.Background(), key)
	return err
}

//create new lease
func (e *Etcd) Lease(duration int64) (leaseResp *clientv3.LeaseGrantResponse, err error) {
	lease := clientv3.NewLease(e.client)
	leaseResp, err = lease.Grant(context.Background(), duration)
	return leaseResp, err
}

//keep lease alive once
func (e *Etcd) KeepLeaseAliveOnce(leaseResp *clientv3.LeaseGrantResponse) (err error) {
	lease := clientv3.NewLease(e.client)
	_, err = lease.KeepAliveOnce(context.Background(), leaseResp.ID)
	return err
}

//close conn
func (e *Etcd) Close() {
	e.client.Close()
}
