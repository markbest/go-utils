package redis

import (
	"encoding/json"
)

type hash struct{}

var Hash = &hash{}

func (h *hash) HSet(key, field string, value interface{}) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HSET", key, field, value)
	return
}

func (h *hash) HGet(key, field string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HGET", key, field)
	return
}

func (h *hash) HMSet(key string, m map[string]map[string][][]string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	for k, v := range m {
		bytes, _ := json.Marshal(&v)
		rc.Do("HMSET", key, k, bytes)
	}
	return
}

func (h *hash) HGetAllValues(key string, clear bool) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HVALS", key)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

func (h *hash) HGetAll(key string, clear bool) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HGETALL", key)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

func (h *hash) HKeys(key string, clear bool) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HKEYS", key)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

func (h *hash) HLen(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HLEN", key)
	return
}

func (h *hash) HDel(key, field string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("HDEL", key, field)
	return
}
