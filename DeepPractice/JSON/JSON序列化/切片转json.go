package main

import (
	"encoding/json"
	"fmt"
)
func main()  {
	dataSlice := make([]map[string]interface{},0)
	datamap1:= make(map[string]interface{})
	datamap1["name"] = "mark"
	datamap1["hobby"] = []string{"篮球","爬山"}
	datamap2:= make(map[string]interface{})
	datamap2["name"] = "markss"
	datamap2["hobby"] = []string{"足球","爬山"}

	dataSlice = append(dataSlice, datamap1,datamap2)
	bytes,err := json.Marshal(dataSlice)
	if err != nil {
		fmt.Println("序列化失败：",err)
		return
	}
	fmt.Println(string(bytes))
}
