package gonoc

import (
	"github.com/garyburd/redigo/redis"
	"time"
	//"reflect"
	//"fmt"
)

var RedisDb RedisPool

type RedisPool struct {
	Pool *redis.Pool
} 

func (rp *RedisPool) CreateRedisPool(server, password string) *redis.Pool {
	if server == "" {
		server = ":6379"
	}
	pool := &redis.Pool{
		MaxIdle :     20,
		MaxActive : 200,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				CheckError(err)
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					CheckError(err)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
	// mq := &MessageQueue{}
	// mq.Pool = pool
	// MQ = mq
	// db := &RedisPool{}
	// db.Pool = pool
	// RedisDb = db
}

func (rp *RedisPool) Lpush(key string, value string) error {
	c := rp.Pool.Get()
	defer c.Close()
	_, err := c.Do("lpush", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RedisPool) Rpush(key string, value string) error {
	c := rp.Pool.Get()
	defer c.Close()
	_, err := c.Do("rpush", key, value)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RedisPool) Lrange(key string, start int, end int) ([]interface{}, error) {
	c := rp.Pool.Get()
	defer c.Close()
	value, err := redis.Values(c.Do("lrange", key, start, end))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (rp *RedisPool) Lpop(key string) (string, error) {
	c := rp.Pool.Get()
	defer c.Close()
	value, err := c.Do("lpop", key)
	if err != nil {
		CheckError(err)
	}
	if value == nil {
		return "", err
	} else {
		return string(value.([]byte)), err
	}
}

func (rp *RedisPool) Lrem(key string, value string) error {
	c := rp.Pool.Get()
	defer c.Close()
	_, err := c.Do("lrem", key, 0, value)
	if err != nil {
		return err
	}
	return nil
}