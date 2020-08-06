package main

import (
	"fmt"
	"net"
	"os"
)
func main()  {
	//创建用于监听的socket
	listener,err := net.Listen("tcp","127.0.0.1:8008")
	if err != nil {
		fmt.Println("listen错误",err)
		return
	}
	defer listener.Close()

	//阻塞监听
	conn,err := listener.Accept()
	if err != nil {
		fmt.Println("Accept错误",err)
		return
	}
	defer conn.Close()

	//获取文件名
	buf := make([]byte,4096)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read错误",err)
		return
	}
	fileName := string(buf[:n])
	//回写OK给发送端
	conn.Write([]byte("ok"))
	//获取文件内容
	recevFile(conn,fileName)
}

func recevFile(conn net.Conn,fileName string)  {
	//按照文件名创建新文件
	f,err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create错误",err)
		return
	}
	defer f.Close()

	//从网络中读数据写入本地文件
	buf := make([]byte,4096)
	for {
		n,_ := conn.Read(buf)
		if n == 0{
			fmt.Println("文件读取完毕")
			return
		}
		//写入本地文件
		f.Write(buf[:n])
	}
}

