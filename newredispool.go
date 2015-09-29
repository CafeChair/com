package com
import (
	"github.com/garyburd/redigo/redis"
)

/*
*redis.Pool struct{
	Dial func() (Conn, error)
	TestOnBorrow func(c Conn, t time.Time) error
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
	...
}
*/
func NewRedisPool(address string) *redis.Pool {
	RedisPool := &redis.Pool{
		// Maximum number of idle connections in the pool
		MaxIdle:10,
		//Dial is for creating and configure connection
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp",address)
			if err != nil {
				return nil,err
			}
			return conn,err
		},
	}
	return RedisPool
}
