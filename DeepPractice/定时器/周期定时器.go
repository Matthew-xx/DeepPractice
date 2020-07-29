package main

import (
	"fmt"
	"time"
)

func main()  {
	quit := make(chan bool)
	myticker := time.NewTicker(time.Second)

	i := 0
	go func() {
		for {
			nowtime := <- myticker.C
			i++
			fmt.Println("nowtime:",nowtime)
			if i == 3 {
				quit <- true
			}
		}
	}()

	<- quit   //循环获取mytime.c直到获取true，否则一致阻塞
}
