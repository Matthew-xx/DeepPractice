package main

import (
	"fmt"
)

type OrderInfo struct {
	Id int
}
func Producer2(out chan <- OrderInfo)  {
	for i:=0;i<10;i++ {
		order := OrderInfo{Id:i+1}   //生成订单
		out <- order
	}
	close(out)
}

func Consumer2(in <- chan OrderInfo)  {
	for order := range in{
		fmt.Println("订单id为：",order.Id)  //处理订单
	}
}

func main()  {
	ch := make(chan OrderInfo)  //可缓存（异步）也可不缓存（同步），看需求
	go Producer2(ch)   //子go程，生成  传只写channel
	Consumer2(ch)  //主go程 消费  传只读channel
}