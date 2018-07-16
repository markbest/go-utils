package redis

type list struct{}

var List = &list{}

func (l *list) LPush(key, value string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("LPUSH", key, value)
	return
}

func (l *list) LGetAll(key string, clear bool) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("LRANGE", key, 0, -1)
	if clear {
		rc.Do("DEL", key)
	}
	return
}

func (l *list) LREM(key string, value string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("LREM", key, 0, value)
	return
}

func (l *list) LLen(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("LLEN", key)
	return
}

func (l *list) RPOP(key string) (reply interface{}, err error) {
	rc := RedisConnPool.Get()
	defer rc.Close()

	reply, err = rc.Do("RPOP", key)
	return
}
