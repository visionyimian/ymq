package mq

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//RedisPool redis连接池
type RedisPool struct {
	Pool *redis.Pool
}

//Redis redis连接
var Redis *RedisPool

func init() {
	Redis = &RedisPool{}
}

//InitPool 初始化连接池
func (r *RedisPool) InitPool() {
	host := "192.168.6.44"
	port := "16379"
	r.Pool = &redis.Pool{
		MaxIdle:     2,   //2个空闲
		MaxActive:   100, //最大数量
		IdleTimeout: 300 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port, redis.DialPassword(""))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
