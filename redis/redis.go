package redis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisConnPool *redis.Pool

func Init(addr, password string, maxIdle, maxActive, db int) {
	RedisConnPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialPassword(password), redis.DialDatabase(db))
			if err != nil {
				log.Fatalf("[Redis] 建立连接失败:原因:%v", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: PingRedis,
	}
	RedisConnPool.Get()
}

func CheckPoolStat() bool {
	conn := RedisConnPool.Get()
	defer conn.Close()

	_, err := conn.Do("ping")
	if err != nil {
		log.Fatalf("[Redis] 心跳ping失败 %v", err)
		return false
	}
	return true
}

func PingRedis(c redis.Conn, t time.Time) error {
	if time.Since(t) < time.Minute {
		return nil
	}

	_, err := c.Do("ping")
	if err != nil {
		log.Fatalf("[Redis] 心跳ping失败 %v", err)
		return err
	}
	return nil
}
