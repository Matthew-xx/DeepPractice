package main

import (
	"fmt"
	"time"
)

func main()  {
	//三种定时器
	//1、sleep
	time.Sleep(time.Second)

	//2、Timer.C
	mytimer := time.NewTimer(time.Second)  //创建定时器，指定定时时长
	mytimer.Reset(15*time.Second)         //重置定时时长
	notime := <- mytimer.C                 //定时器满，系统自动写入当下时间
	fmt.Println("当下时间",notime)
	//mytimer.Stop()                         //停止定时器

	//3、time.After
	nowtime := <- time.After(time.Second)
	fmt.Println("当下时间",nowtime)
}
