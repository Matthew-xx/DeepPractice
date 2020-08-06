package main

import (
	"fmt"
	"net"
)

func main()  {
	listener,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		return
	}
	defer listener.Close()

	conn,_ := listener.Accept()
	defer conn.Close()

	buf := make([]byte,4096)
	n,_ := conn.Read(buf)
	if n == 0 {
		return
	}

	fmt.Printf("**\n%s**\n",string(buf[:n]))
}

//
//**
//GET / HTTP/1.1
//Host: 127.0.0.1:8000
//Connection: keep-alive
//Cache-Control: max-age=0
//Upgrade-Insecure-Requests: 1
//User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
//Accept-Encoding: gzip, deflate, br
//Accept-Language: zh-CN,zh;q=0.9
//
//**
