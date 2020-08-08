
package main

import (
"encoding/json"
"fmt"
"os"
)

func main()  {
	type Person struct {
		Name string
		Age int
		Sex bool
		Hobby []string
	}
	srcFile,_ := os.Open("F:/Software/go_path/my_pro/DeepPractice/JSON/JSON编码映射/切片.json")
	defer srcFile.Close()
	people := make([]Person,0)
	decoder := json.NewDecoder(srcFile)
	err := decoder.Decode(&people)
	if err != nil {
		fmt.Println("解码json文件失败：",err)
		return
	}
	fmt.Println("解码成功:",people)
}
