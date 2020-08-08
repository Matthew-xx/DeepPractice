package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main()  {
	srcFile,_ := os.Open("F:/Software/go_path/my_pro/DeepPractice/JSON/JSON编码映射/测试.json")
	defer srcFile.Close()
	data := make(map[string]interface{})
	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("解码json文件失败：",err)
		return
	}
	fmt.Println("解码成功:",data)
}