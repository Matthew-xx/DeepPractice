package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	data := make(map[string]interface{})
	data["name"] = "mark"
	data["age"] = 23
	data["sex"] = true
	data["hobby"] = []string{"篮球","爬山"}

	bytes,err := json.Marshal(data)
	if err != nil {
		fmt.Println("序列化失败：",err)
		return
	}
	fmt.Println(string(bytes))
}
