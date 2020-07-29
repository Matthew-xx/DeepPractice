package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func producer9(out chan <- int,idx int)  {
	for  {
		cond.L.Lock()  //先加锁
		for len(out) == 5 {  //判断缓冲区是否满
			cond.Wait()
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("生产者%d ,生产：%d\n",idx,num)
		cond.L.Unlock()   //访问公共区结束，解锁
		cond.Signal()  //唤醒阻塞在条件变量上的消费者
		time.Sleep(time.Millisecond*200)
	}
}

func consumer9(in <-chan int,idx int)  {
	for  {
		cond.L.Lock()  //条件变量对应互斥锁加锁
		for len(in) == 0 {  //缓冲区为空，等待生产者生成
			cond.Wait()   //挂起当前协程，等待条件变量满足
		}
		num := <- in   //将channel中的数据读走
		fmt.Printf("****消费者%d, 消费：%d\n",idx,num)
		cond.L.Unlock()  //消费结束，解锁互斥锁
		cond.Signal()   //唤醒阻塞的生产者
		time.Sleep(time.Millisecond*200)
	}
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)

	pruduct := make(chan int,5)  //公共区，channel模拟
	cond.L = new(sync.Mutex)   //创建互斥锁和条件变量

	for i:=0;i<5 ;i++  {
		go producer9(pruduct,i+1)
	}
	for i:=0;i<3 ;i++  {
		go consumer9(pruduct,i+1)
	}
	<-quit   //主协程阻塞，不结束
}
