package utils

import (
	"testing"
	"bytes"
)

var (
	host     = "127.0.0.1"
	port     = 22
	user       = "root"
	password = "password"
)

func TestNewSShClient(t *testing.T) {
	client, err := NewSShClient(host, int64(port), user, password)
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
}

func TestSSHClient_Commands(t *testing.T) {
	client, _ := NewSShClient(host, int64(port), user, password)
	defer client.Close()

	var outPut bytes.Buffer
	commands := []string{"ls -l"}
	err := client.Commands(commands, outPut)
	if err != nil {
		t.Error(err)
	}
	t.Log(outPut.String())
}