package utils

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

var (
	mongodbHost     = "127.0.0.1"
	mongodbPort     = "27017"
	mongodbSource   = "admin"
	mongodbUsername = ""
	mongodbPassword = ""
)

type Book struct {
	Title  string
	Img    string
	Author string
	Sell   string
	Url    string
}

func TestMongodb_Connect(t *testing.T) {
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
	defer m.Close()
}

func TestMongodb_One(t *testing.T) {
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
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
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
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
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
	defer m.Close()

	rs := m.DB("crawler").AllCollections()
	t.Log(rs)
}

func TestMongodb_Insert(t *testing.T) {
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
	defer m.Close()

	rs := &Book{"title", "img", "author", "sell", "url"}
	err := m.DB("crawler").Collection("book").Insert(rs)
	if err != nil {
		t.Error(err)
	}
}

func TestMongodb_Update(t *testing.T) {
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
	defer m.Close()

	err := m.DB("crawler").Collection("book").Update(bson.M{"title": "title"}, bson.M{"title": "title1"})
	if err != nil {
		t.Error(err)
	}
}

func TestMongodb_Remove(t *testing.T) {
	m := NewMongodb(mongodbHost, mongodbPort, mongodbSource, mongodbUsername, mongodbPassword)
	defer m.Close()

	err := m.DB("crawler").Collection("book").Remove(bson.M{"title": "title"})
	if err != nil {
		t.Error(err)
	}
}
