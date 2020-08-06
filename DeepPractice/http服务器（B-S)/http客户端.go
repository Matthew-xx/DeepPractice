package main

import (
	"fmt"
	"net/http"
)

func main()  {
	//获取服务器，应答包内容
	resp,err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get err",err)
		return
	}
	defer resp.Body.Close()

	//简单查看应答包
	fmt.Println("header",resp.Header)
	fmt.Println("Status",resp.Status)
	fmt.Println("Proto",resp.Proto)


	buf := make([]byte,4096)
	var result string
	for {
		n,_ := resp.Body.Read(buf)
		if n==0 {
			fmt.Println("read finish")
			break
		}
		//if err != nil {
		//	fmt.Println("body read err",err)
		//	break
		//	//return   //选return的话，fmt.Printf("|%v|\n",result)还没打印就返回了
		//}
		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n",result)
}
