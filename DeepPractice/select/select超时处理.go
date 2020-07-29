package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for  {
			select {
			case num:= <-ch:
				fmt.Println("num=",num)
			case <- time.After(3*time.Second):  //超时写入
				quit <- true
				return
			}
		}
	}()

	<- quit
	fmt.Println("finish!")
}
