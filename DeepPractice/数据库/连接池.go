package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

//数据库的连接是IO资源，是有限的，避免无度的开辟。
//在池子里放一堆连接，这部分连接都用完后，其他人等待。
//可避免DDOS攻击（分布式攻击）：无意义的访问占满连接，占满IO

func main()  {
	//配置连接池对象
	pool := &redis.Pool{
		MaxIdle:20,  //最大闲置连接数
		MaxActive:0, //最大活动连接数，0=无限
		IdleTimeout:time.Second*100,  //闲置连接的超时时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp","127.0.0.1:6379")
		},//定义拨号获得连接的函数
	}
	defer pool.Close()

	//并发连接
	for i:= 0;i<10;i++ {
		go getConn(pool,i)
	}
	time.Sleep(3*time.Second)  //保持主协程
}

func getConn(pool *redis.Pool,i int)  {
	conn := pool.Get()
	defer conn.Close()

	reply,err := conn.Do("set","conn"+strconv.Itoa(i),i)
	s,_ := redis.String(reply,err)
	fmt.Println(s)
}

