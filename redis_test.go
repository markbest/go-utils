package utils

import "testing"

var (
	redisHost     = "127.0.0.1"
	redisPort     = "6379"
	redisDb       = 1
	redisPassword = ""
)

type Article struct {
	Title  string `redis:"title"`
	Author string `redis:"author"`
	Body   string `redis:"body"`
}

func TestCacheRedis_Set(t *testing.T) {
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	if err := redis.Set("test", "test content", 0); err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_Get(t *testing.T) {
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	v, err := redis.Get("test")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_HSet(t *testing.T) {
	art := Article{"Example", "Gary", "Hello World!"}
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	if err := redis.HSet("article:1", &art, 0); err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_HGet(t *testing.T) {
	var art Article
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	if err := redis.HGet("article:1", &art); err != nil {
		t.Error(err)
	} else {
		t.Log(art)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_IsExist(t *testing.T) {
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	_, err := redis.IsExist("test")
	if err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_Del(t *testing.T) {
	redis := NewRedis(redisHost, redisPort, redisDb, redisPassword)
	err := redis.Del("test")
	if err != nil {
		t.Error(err)
	}

	err = redis.Del("article:1")
	if err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}
