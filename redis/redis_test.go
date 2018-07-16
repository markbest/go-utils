package redis

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

var (
	redisAddr     = "127.0.0.1:6379"
	redisDb       = 1
	redisPassword = ""
)

func TestRedis(t *testing.T) {
	Init(redisAddr, redisPassword, 100, 500, redisDb)

	// set key
	_, err := Key.Set("test", "123456")
	if err != nil {
		t.Error(err.Error())
	}

	// get key
	value, err := redis.String(Key.Get("test"))
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(value)

	// expire key
	Key.Expire("test", 10)
}
