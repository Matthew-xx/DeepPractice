package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:666666@tcp(127.0.0.1:3306)/beeblog?charset=utf8",30)

	//orm.RegisterModel(new(Category),new(Topic),new(Comment))  //每添加新表需先注册

	o := orm.NewOrm()
	o.Begin()  //开始事务
	o.Rollback()  //失败后回滚
	o.Commit()  //全部完成后提交
	/*事务的应用场景很多，例如在一次电商的交易中，只有资金出入表、物流表、交易记录表、订单表、购物车表等一系列表的数据变动全部严丝合缝分毫不差时，才能视为交易成功，此时可以提交事务；
	只要有一点对不上号的地方，本次交易就不能视为成功，所有的数据应全部回滚至交易之前的状态；

	 */

}
