package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Human struct {
	Name string  `db:"name"`
	Age int      `db:"age"`
}

func main()  {
	var cmd string
	for {
		fmt.Println("请输入命令：")
		fmt.Scan(&cmd)

		switch cmd {
		case "getall":
			GetAll()
		case "exit":
			goto OVER
		default:
			fmt.Println("请重新输入!")
		}
	}
	OVER:
	fmt.Println("查询结束")
	
}
func GetAll()  {
	//先从缓存读取
	peoplestr := GetFromRedis()

	if len(peoplestr) ==0 || peoplestr == nil{
		people := GetFromSql()
		//缓存查询结果到redis
		CacheToRedis(&people)  //传地址，可减少拷贝
	}else {
		fmt.Println(peoplestr)
	}

}

func CacheToRedis(people *[]Human)  {
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	conn.Do("del","people")  //清除原有缓存
	for _, human := range *people {
		_,err :=conn.Do("rpush","people",fmt.Sprint(human))
		HandlerErr(err,"缓存失败")
	}
	//设置过期时间
	_,err := conn.Do("expire","people",60)
	HandlerErr(err,"设置超时失败。")
	fmt.Println("缓存成功！")
}

func HandlerErr(err error,why string)  {
	if err != nil {
		fmt.Println(err,why)
		os.Exit(1)
	}
}

func GetFromRedis() (people []string) {
	fmt.Println("从redis读取数据！")
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	defer conn.Close()

	reply,err := conn.Do("lrange","people","0","-1")
	HandlerErr(err,"redis cahe失败")
	people,_ = redis.Strings(reply,err)
	fmt.Println(people)
	return
}

func GetFromSql() (people []Human)  {

	db,_ :=sqlx.Connect("mysql","root:666666@tcp(localhost:3306)/mydb")
	defer db.Close()

	err := db.Select(&people,"select name,age from person")
	if err != nil {
		fmt.Println("查询失败：",err)
		return
	}
	fmt.Println(people)
	return
}