package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	//主动发起连接请求
	//conn,err :=net.Dial("tcp","127.0.0.1:8001")
	conn,err :=net.Dial("tcp","192.168.0.5:8001")
	if err != nil {
		fmt.Println("Dial err:",err)
		return
	}
	defer conn.Close()

	//获取用户输入，将输入数据发送给服务器
	go func() {
		str := make([]byte,4096)
		for  {
			n,err :=os.Stdin.Read(str)
			if err != nil{
				fmt.Println("os read err:",err)
				continue
			}
			//写给服务器
			conn.Write(str[:n])
		}
	}()

	//显示服务器回发的内容
	buf := make([]byte,4096)
	for {
		n,err := conn.Read(buf)
		if n == 0 {
			fmt.Println("退出连接")
			return
		}
		if err != nil {
			fmt.Println("conn.read err:",err)
			return
		}
		fmt.Println("服务器回发：",string(buf[:n]))
	}
}
