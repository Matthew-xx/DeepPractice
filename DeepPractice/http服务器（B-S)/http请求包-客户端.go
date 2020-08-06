package main

import (
	"fmt"
	"net"
)

func main()  {
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		return
	}
	defer conn.Close()

	httpRequest := "GET /home HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n" //请求行+请求头+空行
	conn.Write([]byte(httpRequest))
	buf := make([]byte,4096)
	n,_ := conn.Read(buf)
	if n==0 {
		return
	}
	fmt.Printf("**\n%s**\n",string(buf[:n]))
}


//**
//HTTP/1.1 200 OK
//Date: Thu, 06 Aug 2020 07:09:20 GMT
//Content-Length: 12
//Content-Type: text/plain; charset=utf-8
//
//Hello Google**
