package main

import (
	"fmt"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter,r *http.Request)  {
	OpenSendFile(r.URL.String(),w)
}

func OpenSendFile(fileName string,w http.ResponseWriter)  {
	pathFileName := "C:/Users/mark/Desktop/电商/开发"+fileName
	f,err := os.Open(pathFileName)
	if err != nil{
		fmt.Println("file open err:",err)
		w.Write([]byte("No such file or directory!"))
		return
	}
	defer f.Close()

	buf := make([]byte,4096)
	for {
		n,_ := f.Read(buf)  //从本地文件内容读取
		if n==0 {
			return
		}
		w.Write(buf[:n])
	}
}

func main()  {
	//1、注册回调函数，该回调函数会在服务器被访问时自动被调用
	http.HandleFunc("/",myHandler)
	//2、绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000",nil)
}
