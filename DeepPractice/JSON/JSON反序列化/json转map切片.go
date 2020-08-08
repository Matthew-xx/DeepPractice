package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsonStr := `[{"hobby":["篮球","爬山"],"name":"mark"},{"hobby":["足球","爬山"],"name":"markss"}]`
	jsonBytes := []byte(jsonStr)
	dataSlice := make([]map[string]interface{},0)
	err := json.Unmarshal(jsonBytes,&dataSlice)
	if err != nil {
		fmt.Println("反序列化失败：",err)
		return
	}
	fmt.Println(dataSlice)
}
