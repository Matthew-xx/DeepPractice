package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	Sex bool
	Hobby []string
}

func main()  {
	person := Person{"mark",23,true,[]string{"篮球","爬山"}}
	bytes,err := json.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败：",err)
		return
	}
	fmt.Println(string(bytes))
}

