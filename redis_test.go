package utils

import "testing"

var (
	redis_host     = "127.0.0.1"
	redis_port     = "6379"
	redis_db       = 1
	redis_password = ""
)

type Article struct {
	Title  string `redis:"title"`
	Author string `redis:"author"`
	Body   string `redis:"body"`
}

func TestCacheRedis_Set(t *testing.T) {
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
	if err := redis.Set("test", "test content", 0); err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_Get(t *testing.T) {
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
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
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
	if err := redis.HSet("article:1", &art, 0); err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_HGet(t *testing.T) {
	var art Article
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
	if err := redis.HGet("article:1", &art); err != nil {
		t.Error(err)
	} else {
		t.Log(art)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_IsExist(t *testing.T) {
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
	_, err := redis.IsExist("test")
	if err != nil {
		t.Error(err)
	}
	defer redis.CloseRedis()
}

func TestCacheRedis_Del(t *testing.T) {
	redis := NewRedis(redis_host, redis_port, redis_db, redis_password)
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
