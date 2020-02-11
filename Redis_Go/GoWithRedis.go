package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//Asking connect for a redis Server
	//localhost represent for local IP
	conn,err := redis.Dial("tcp","localhost:6379")
	if err != nil{
		fmt.Println("Dial error:",err)
		return
	}
	_,err = conn.Do("Set","name","liu")
	if err != nil{
		fmt.Println("Set error:",err)
	}

	//Or use redis.Strings to deal with more then one result. This function return a slice
	//Also redis.Int or redis.Float64 is ok
	s,err := redis.String(conn.Do("Get","name"))
	if err != nil{
		fmt.Println("Get error:",err)
	}else{
		fmt.Println("Get from redis:",s)
	}
	//fmt.Println("Connection success:",conn)

}
