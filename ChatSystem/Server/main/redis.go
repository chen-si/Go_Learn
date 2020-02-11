package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//不会自动调用
func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxActive:   maxActive,   //表示和数据库的最大链接数 0表示没有限制
		MaxIdle:     maxIdle,     //最大空闲链接数
		IdleTimeout: idleTimeout, //最大空闲链接
		Dial: func() (redis.Conn, error) { //初始化链接
			return redis.Dial("tcp", address)
		},
	}
}
