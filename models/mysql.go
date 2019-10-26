package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-deiver/mysql"
	"BeegoTest/models/table"
)

func init()  {
	// 注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/BeegoTest")

	// 注册模型
	orm.RegisterModel(new(Users))

	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", true, true)
}

