package main

import (
	"fmt"
	"net/http"
)

func main()  {
	//1、注册回调函数，该回调函数会在服务器被访问时自动被调用
	http.HandleFunc("/home",handler)
	//2、绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000",nil) //传 nil 时会调用默认自带的函数（通常传空
}

//回调函数（必须是(w http.ResponseWriter,r *http.Request)
func handler(w http.ResponseWriter,r *http.Request)  {
	//w:写给客户端（浏览器）的数据
	w.Write([]byte("Hello Google"))
	//r:从客户端（浏览器）读到的数据
	fmt.Println("Header:",r.Header)
	fmt.Println("Method:",r.Method)
	fmt.Println("Body:",r.Body)

}