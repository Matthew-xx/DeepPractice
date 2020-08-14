package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)
func SHandlerErr(err error,when string)  {
	if err != nil {
		fmt.Println(when,err)
		os.Exit(1)
	}
}

func main() {
	//建立tcp监听
	listener,err :=net.Listen("tcp","127.0.0.1:8088")
	SHandlerErr(err,"net.listener error")
	defer func() {
		listener.Close()
		fmt.Println("服务端正常退出")
	}()
	//接受客户端请求，建立conn
	connSev,err := listener.Accept()
	SHandlerErr(err,"listenerAccept error")
	defer func() {
		connSev.Close()
		fmt.Printf("已断开与客户端%v的连接\n",connSev.RemoteAddr())
	}()

	sFile,_ := os.OpenFile(`F:\Software\go_path\my_pro\DeepPractice\文件上传\Sfile\recv.png`,os.O_CREATE | os.O_WRONLY | os.O_TRUNC,666)
	writer := bufio.NewWriter(sFile)  //缓冲器写入
	defer sFile.Close()
	buffSev := make([]byte,100)


	for {
		//接收客户端上传的文件
		n,err := connSev.Read(buffSev)
		SHandlerErr(err,"connRead error")

		//写入服务端本地文件
		sFile.Write(buffSev[:n])

		fmt.Println("成功写入文件")

		//代表文件已读完（没填满一个缓冲区
		if n <100 {
			writer.Flush()
			//回复客户端
			connSev.Write([]byte("文件上传成功"))
			break
		}
	}


}
