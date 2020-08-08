package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main()  {
	data := make(map[string]interface{})
	data["name"] = "mark"
	data["age"] = 23
	data["sex"] = true
	data["hobby"] = []string{"篮球","爬山"}
	jsonFile,_ := os.OpenFile("F:/Software/go_path/my_pro/DeepPractice/JSON/JSON编码映射/测试.json",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,666)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err := encoder.Encode(data)  //将数据编码到json文件
	if err != nil {
		fmt.Println("编码到json文件失败：",err)
		return
	}
	fmt.Println("编码成功！")
}
