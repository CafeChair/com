package com
import (
	"github.com/garyburd/redigo/redis"
)

var (
	MAX_POOL_SIZE = 20
	RedisPool chan redis.Conn
)

func PutInRedisPool(conn redis.Conn) {
	if RedisPool == nil {
		RedisPool = make(chan redis.Conn, MAX_POOL_SIZE)
	}
	if len(RedisPool) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	RedisPool <- conn
}

func InitRedisConn(network, address string) redis.Conn {
	if len(RedisPool) == 0 {
		RedisPool = make(chan redis.Conn, MAX_POOL_SIZE)
		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				conn,err := redis.Dial(network, address)
				if err != nil {
					panic(err)
				}
				PutInRedisPool(conn)
			}
		}()
	}
	return <- RedisPool
}
