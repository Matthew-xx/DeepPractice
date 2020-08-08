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

	p1 := Person{"mark",24,true,[]string{"篮球，足球"}}
	p2 := Person{"ma",22,true,[]string{"羽毛球，足球"}}
	p3 := Person{"ke",25,true,[]string{"篮球，爬山"}}
	people := make([]Person,0)
	people = append(people,p1,p2,p3)

	jsonFile,_ := os.OpenFile("F:/Software/go_path/my_pro/DeepPractice/JSON/JSON编码映射/测试.json",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,666)
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err := encoder.Encode(people)
	if err != nil {
		fmt.Println("编码到json文件失败：",err)
		return
	}
	fmt.Println("编码成功！")
}
