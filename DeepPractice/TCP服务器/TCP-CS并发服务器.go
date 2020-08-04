package main

import (
	"fmt"
	"net"
	"strings"
)

func main()  {
	//创建监听套接字
	//listener,err := net.Listen("tcp","127.0.0.1:8001")
	listener,err := net.Listen("tcp","192.168.0.5:8001")
	if err!= nil {
		fmt.Println("net.listener err:",err)
		return
	}
	defer listener.Close()

	for  {
		fmt.Println("服务器等待连接...")
		//监听客户端连接请求(循环监听多个客户端
		conn,err := listener.Accept()
		if err!= nil {
			fmt.Println("Accept err:",err)
			return
		}

		//具体完成服务器和客户端的数据通信(go程，连接一个客户端开一个go程去处理，
		// 而不用全部直接去请求服务器，如果是全部直接去请求服务器，那么需要排队等待处理
		go HandlerConnect(conn)
	}

}

func HandlerConnect(conn net.Conn)  {
	defer conn.Close()
	
	addr := conn.RemoteAddr()
	fmt.Println(addr,"客户端成功连接")
	//循环读取客户端发送的数据
	buf := make([]byte,4096)
	for  {
		//for循环避免一次性操作,直到客户端关闭
		n,err := conn.Read(buf)
		if string(buf[:n]) == "exit\r\n" || string(buf[:n]) == "exit\n"{
			fmt.Println("客户端退出")
			return
		}
		if n == 0 {
			fmt.Println("客户端已断开")
			return
		}
		if err != nil {
			fmt.Println("Read err:",err)
			return
		}
		fmt.Println("服务器读取数据：",string(buf[:n])) //使用数据
		//处理数据,回发给客户端
		_,err = conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		if err!= nil {
			fmt.Println("Write err:",err)
			return
		}
	}
}
