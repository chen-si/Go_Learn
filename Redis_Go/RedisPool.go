// package main

// import (
// 	"fmt"
// 	"github.com/garyburd/redigo/redis"
// )

// var pool *redis.Pool

// func init() {
// 	pool = &redis.Pool{
// 		MaxActive:   0,
// 		MaxIdle:     8,
// 		IdleTimeout: 300,
// 		Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", "localhost:6379")
// 		},
// 	}
// }
// func main() {
// 	//pool.Close()
// 	conn := pool.Get()
// 	defer pool.Close()
// 	_, err := conn.Do("HMSet", "User2", "name", "liu", "age", 20, "sex", "man")
// 	if err != nil {
// 		fmt.Println("HMSet error:", err)
// 	}
// 	s, err := redis.Strings(conn.Do("HMGet", "User2", "name", "age", "sex"))
// 	if err != nil {
// 		fmt.Println("HMGet error:", err)
// 	} else {
// 		fmt.Println(s)
// 	}
// }
