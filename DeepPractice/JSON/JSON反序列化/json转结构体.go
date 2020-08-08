package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsonStr := `{"age":23,"hobby":["篮球","爬山"],"name":"mark","sex":true}`
	type Person struct {
		Name string
		Age int
		Sex bool
		Hobby []string
	}
	jsonBytes := []byte(jsonStr)
	person := new(Person)
	err := json.Unmarshal(jsonBytes,person)
	if err != nil {
		fmt.Println("反序列化失败：",err)
		return
	}
	fmt.Println(*person)
}
