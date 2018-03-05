package utils

import "testing"

var (
	etcd_servers = []string{"127.0.0.1:2379"}
)

func TestEtcd_Put(t *testing.T) {
	client := NewEtcdConn(etcd_servers)
	defer client.Close()

	if err := client.Put("/test/1", "test1"); err != nil {
		t.Error(err)
	}

	if err := client.Put("/test/2", "test2"); err != nil {
		t.Error(err)
	}
}

func TestEtcd_Get(t *testing.T) {
	client := NewEtcdConn(etcd_servers)
	defer client.Close()

	if value, err := client.Get("/test/1"); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}

func TestEtcd_GetKeysByPrefix(t *testing.T) {
	client := NewEtcdConn(etcd_servers)
	defer client.Close()

	if values, err := client.GetKeysByPrefix("/test"); err != nil {
		t.Error(err)
	} else {
		t.Log(values)
	}
}

func TestEtcd_Del(t *testing.T) {
	client := NewEtcdConn(etcd_servers)
	defer client.Close()

	if err := client.Del("/test/1"); err != nil {
		t.Error(err)
	}

	if err := client.Del("/test/2"); err != nil {
		t.Error(err)
	}
}
