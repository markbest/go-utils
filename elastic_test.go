package utils

import (
	"log"
	"os"
	"testing"
)

var (
	esHost = "http://127.0.0.1:9300"
)

func TestES_CreateIndex(t *testing.T) {
	es := NewES(esHost, log.New(os.Stdout, "", log.LstdFlags))
	if err := es.CreateIndex("test"); err != nil {
		t.Error(err)
	}
}

func TestES_DeleteIndex(t *testing.T) {
	es := NewES(esHost, log.New(os.Stdout, "", log.LstdFlags))
	if err := es.DeleteIndex("test"); err != nil {
		t.Error(err)
	}
}
