package utils

import "testing"

var (
	zk_servers = []string{"118.31.41.54"}
)

func TestZookeeper_Create(t *testing.T) {
	zk := NewZKConn(zk_servers)
	defer zk.Close()

	if err := zk.Create("/test", []byte("test"), 0); err != nil {
		t.Error(err)
	}
}

func TestZookeeper_Get(t *testing.T) {
	zk := NewZKConn(zk_servers)
	defer zk.Close()

	if s, _, err := zk.Get("/test"); err != nil {
		t.Error(err)
	} else {
		t.Log(string(s))
	}
}

func TestZookeeper_Update(t *testing.T) {
	zk := NewZKConn(zk_servers)
	defer zk.Close()

	if err := zk.Update("/test", []byte("test_update"), 2); err != nil {
		t.Error(err)
	} else {
		t.Log("update success.")
	}
}

func TestZookeeper_Exist(t *testing.T) {
	zk := NewZKConn(zk_servers)
	defer zk.Close()

	if flag, err := zk.Exist("/test"); err != nil {
		t.Error(err)
	} else {
		if flag {
			t.Log("Path is exist.")
		} else {
			t.Log("Path is not exist.")
		}
	}
}

func TestZookeeper_Delete(t *testing.T) {
	zk := NewZKConn(zk_servers)
	defer zk.Close()

	if err := zk.Delete("/test", 3); err != nil {
		t.Error(err)
	}
}
