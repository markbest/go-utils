package redis

type key struct{}

var Key = &key{}

func (k *key) Set(key string, value interface{}) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("SET", key, value)
	return
}

func (k *key) Expire(key string, seconds uint) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	rc.Do("EXPIRE", key, seconds)
}

func (k *key) Append(key, value string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("APPEND", key, value)
	return
}

func (k *key) Get(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("GET", key)
	return
}

func (k *key) Decr(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("DECR", key)
	return
}

func (k *key) DecrBy(key string, num int) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("DECRBY", key, num)
	return
}

func (k *key) Incr(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("INCR", key)
	return
}

func (k *key) IncreBy(key string, num int) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("INCRBY", key, num)
	return
}

func (k *key) Del(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("DEL", key)
	return
}

func (k *key) Clear(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()
	reply, err = rc.Do("SET", key, "")
	return
}
