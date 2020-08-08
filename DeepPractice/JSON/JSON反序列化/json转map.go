package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsonStr := `{"age":23,"hobby":["篮球","爬山"],"name":"mark","sex":true}`
	jsonBytes := []byte(jsonStr)
	date := make(map[string]interface{})
	err := json.Unmarshal(jsonBytes,&date)
	if err != nil {
		fmt.Println("反序列化失败：",err)
		return
	}
	fmt.Println(date)
}
