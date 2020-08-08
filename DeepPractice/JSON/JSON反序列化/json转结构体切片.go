package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	type Person struct {
		Name string
		Age int
		Sex bool
		Hobby []string
	}
	jsonStr := `[{"hobby":["篮球","爬山"],"name":"mark"},{"hobby":["足球","爬山"],"name":"markss"}]`
	jsonBytes := []byte(jsonStr)
	persons := make([]Person,0)
	err := json.Unmarshal(jsonBytes,&persons)
	if err != nil {
		fmt.Println("反序列化失败：",err)
		return
	}
	fmt.Println(persons)
}
