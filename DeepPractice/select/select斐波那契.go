package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <- chan int,quit <- chan bool)  {
	for  {
		select {
		case num := <- ch:
			fmt.Print(num," ")
		case <- quit:
			runtime.Goexit()   //等效于return
		}
	}
}

func main()  {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch,quit)

	x,y := 1,1
	for i:= 0;i<10;i++ {
		ch <- x
		x,y = y,x+y
	}
	quit <- true
}
