package utils

import (
	"testing"
)

var (
	dbHost     = "127.0.0.1"
	dbPort     = 3306
	dbUser     = "root"
	dbPassword = ""
	dbDatabase = "test"
	dbTable    = "test"
)

func TestDB_Connect(t *testing.T) {
	db, err := NewDB(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestDB_Insert(t *testing.T) {
	db, err := NewDB(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	values := make(map[string]interface{})
	values["name"] = "mark"
	values["age"] = 29
	values["sex"] = 1
	values["province"] = "陕西"
	if _, err = db.Insert(dbTable, values); err != nil {
		t.Error(err)
	}
}

func TestDB_Update(t *testing.T) {
	db, err := NewDB(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	conditions := make(map[string]interface{})
	conditions["name"] = "mark"
	values := make(map[string]interface{})
	values["age"] = 18

	if _, err = db.Update(dbTable, conditions, values); err != nil {
		t.Error(err)
	}
}

func TestDB_Select(t *testing.T) {
	db, err := NewDB(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	conditions := make(map[string]interface{})
	conditions["name"] = "mark"

	rs, err := db.Select(dbTable, conditions)
	if err != nil {
		t.Error(err)
	}

	t.Log(rs)
}

func TestDB_Delete(t *testing.T) {
	db, err := NewDB(dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	conditions := make(map[string]interface{})
	conditions["name"] = "mark"

	if _, err = db.Delete(dbTable, conditions); err != nil {
		t.Error(err)
	}
}
