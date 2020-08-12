package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	//conn,err := redis.Dial("tcp","192.168.99.100:6379")
	if err != nil {
		return
	}
	defer conn.Close()
	reply,_ := conn.Do("name","1236")
	fmt.Println(reply)
}
