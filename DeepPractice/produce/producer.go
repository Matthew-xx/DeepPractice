package main

import (
	"fmt"
)

func Producer(out chan <- int)  {
	for i:=0;i<10;i++ {
		fmt.Println("生产者生成：",i*i)
		out <- i*i
	}
	close(out)
}

func Consumer(in <- chan int)  {
	for num := range in{
		fmt.Println("消费者拿到：",num)
	}
}

func main()  {
	ch := make(chan int,5)  //可缓存（异步）也可不缓存（同步），看需求
	go Producer(ch)   //子go程，生成
	Consumer(ch)  //主go程 消费
}