package redigo

import (
	"encoding/json"
	"ginLearn/pkg/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPools *redis.Pool

func Init() error {
	RedisPools = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: config.RedisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			pool, err := redis.Dial("tcp", config.RedisConfig.Host)
			if err != nil {
				return nil, err
			}
			if config.RedisConfig.Password != "" {
				if _, err := pool.Do("AUTH", config.RedisConfig.Password); err != nil {
					pool.Close()
					return nil, err
				}
			}
			return pool, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
func Set(key string, data interface{}) error {
	conn := RedisPools.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
func SetStr(key string, value string) error {
	conn := RedisPools.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
func SetStrWithExpire(key string, value string, second int) error {
	conn := RedisPools.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value, "EX", second)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := RedisPools.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func GetStr(key string) (string, error) {
	conn := RedisPools.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisPools.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisPools.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
