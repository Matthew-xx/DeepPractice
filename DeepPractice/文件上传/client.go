package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func CHandlerErr(err error,when string)  {
	if err != nil {
		fmt.Println(when,err)
	}
}
func main() {
	//连接
	conn,err := net.Dial("tcp","127.0.0.1:8088")
	CHandlerErr(err,"net.dial error")
	defer func() {
		conn.Close()
		fmt.Println("客户端正常退出")
	}()

	//读取文件并写入(小文件一次性读取
	bytes,err := ioutil.ReadFile(`F:\Software\go_path\my_pro\DeepPractice\文件上传\Cfile\http.png`)
	CHandlerErr(err,"ioRead error")
	_,err = conn.Write(bytes)
	CHandlerErr(err,"connWrite error")

	//拿到客户端返回字段
	buff := make([]byte,1024)
	n,err := conn.Read(buff)
	CHandlerErr(err,"connRead error")
	reply := buff[:n]
	fmt.Println("服务端：",string(reply))
}
