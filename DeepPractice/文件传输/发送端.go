package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main()  {
	list := os.Args  //获取命令行参数
	if len(list) != 2{
		fmt.Println("格式错误")
		return
	}
	filePath := list[1]  //文件路径
	fileInfo ,err := os.Stat(filePath)  //文件信息
	if err != nil {
		fmt.Println("文件读取错误",err)
		return
	}
	fileName := fileInfo.Name()
	//主动发起连接请求
	conn,err := net.Dial("tcp","127.0.0.1:8008")
	if err != nil {
		fmt.Println("连接请求失败",err)
		return
	}
	defer conn.Close()

	//发送文件名给接收端
	_,err = conn.Write([]byte(fileName))

	//读取服务器回发的信息
	buf := make([]byte,4096)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("回执读取错误",err)
		return
	}
	if string(buf[:n]) == "ok"{
		//写文件内容给服务器
		sendFile(conn,filePath)
	}
}

func sendFile(conn net.Conn,filePath string)  {
	//只读打开文件
	f,err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开错误",err)
		return
	}
	defer f.Close()
	//从本地文件读取数据写给网络接收端，读多少写多少
	buf := make([]byte,4096)
	for {
		n,err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			}else {
				fmt.Println("f文件读取错误",err)
			}
			return
		}
		//写到网络中
		_,err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write错误",err)
			return
		}
	}
}
