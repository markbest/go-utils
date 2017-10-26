package utils

import (
	"testing"
)

var (
	mongodb_host     = "127.0.0.1"
	mongodb_port     = "27017"
	mongodb_source   = "admin"
	mongodb_username      = ""
	mongodb_password = ""
)
type Book struct {
	Title  string
	Img    string
	Author string
	Sell   string
	Url    string
}

func TestMongodb_Connect(t *testing.T) {
	m := NewMongodb(mongodb_host, mongodb_port, mongodb_source, mongodb_username, mongodb_password)
	defer m.Close()
}

func TestMongodb_One(t *testing.T) {
	m := NewMongodb(mongodb_host, mongodb_port, mongodb_source, mongodb_username, mongodb_password)
	defer m.Close()

	rs := &Book{}
	err := m.DB("crawler").Collection("book").Where(nil).One(rs)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(rs)
	}
}

func TestMongodb_All(t *testing.T) {
	m := NewMongodb(mongodb_host, mongodb_port, mongodb_source, mongodb_username, mongodb_password)
	defer m.Close()

	rs := &[]Book{}
	err := m.DB("crawler").Collection("book").Where(nil).Limit(1).All(rs)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(rs)
	}
}

func TestMongodb_AllCollections(t *testing.T) {
	m := NewMongodb(mongodb_host, mongodb_port, mongodb_source, mongodb_username, mongodb_password)
	defer m.Close()

	rs := m.DB("crawler").AllCollections()
	t.Log(rs)
}