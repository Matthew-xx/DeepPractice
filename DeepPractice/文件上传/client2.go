package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func CHandlerErr2(err error,when string)  {
	if err != nil {
		fmt.Println(when,err)
	}
}
func main() {
	//连接
	conn,err := net.Dial("tcp","127.0.0.1:8088")
	CHandlerErr2(err,"net.dial error")
	defer func() {
		conn.Close()
		fmt.Println("客户端正常退出")
	}()

	buf := make([]byte,100)
	srcFile,_ := os.Open(`F:\Software\go_path\my_pro\DeepPractice\文件上传\Cfile\http.png`)
	reader := bufio.NewReader(srcFile)
	for {
		n,err := reader.Read(buf)
		_,_ = conn.Write(buf[:n])

		if err == io.EOF {
			fmt.Println("文件上传完毕")
			break
		}
	}


	//拿到客户端返回字段
	buff := make([]byte,1024)
	n,err := conn.Read(buff)
	CHandlerErr2(err,"connRead error")
	reply := buff[:n]
	fmt.Println("服务端：",string(reply))
}
