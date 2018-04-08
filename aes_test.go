package utils

import (
	"encoding/hex"
	"testing"
)

const (
	key  = "3195c71b9af2d5a755eae16562f41142"
	iv   = "9eb8b26e9f3bcd130d1c23cdff8740ef"
	data = "Hello World!"
)

func TestEncrypt(t *testing.T) {
	encryptKey, _ := hex.DecodeString(key)
	encryptIV, _ := hex.DecodeString(iv)
	enStr, err := Encrypt(data, encryptKey, encryptIV)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(enStr)
	}
}

func TestDecrypt(t *testing.T) {
	encryptKey, _ := hex.DecodeString(key)
	encryptIV, _ := hex.DecodeString(iv)
	enStr, err := Encrypt(data, encryptKey, encryptIV)
	if err != nil {
		t.Error(err)
	}
	deStr, err := Decrypt(enStr, encryptKey, encryptIV)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(deStr)
	}
}
