package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

func InitPool(host string, port int, password string, db int, timeout int, maxIdle int, maxActive int, idleTimeout int) {

	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port),
				redis.DialPassword(password),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(time.Duration(timeout)*time.Second),
				redis.DialReadTimeout(time.Duration(timeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(timeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func Get() redis.Conn {
	if redisPool == nil {
		return nil
	}

	return redisPool.Get()
}

func DoStrSet(key string, value interface{}) error {
	conn := Get()
	if conn == nil {
		return fmt.Errorf("conn is nil")
	}
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

func DoStrGet(key string) (string, error) {
	conn := Get()
	if conn == nil {
		return "", fmt.Errorf("conn is nil")
	}
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

func DoExpire(key string, expire time.Duration) error {
	conn := Get()
	if conn == nil {
		return fmt.Errorf("conn is nil")
	}
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, expire)
	return err
}

func DoExists(key string) (bool, error) {
	conn := Get()
	if conn == nil {
		return false, fmt.Errorf("conn is nil")
	}
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func DoDel(key string) error {
	conn := Get()
	if conn == nil {
		return false, fmt.Errorf("conn is nil")
	}
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}
