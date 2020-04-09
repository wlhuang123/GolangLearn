package controllers

import (
	"fmt"
	"hwl/tool/logs"

	"github.com/astaxie/beego"         // beego的web端
	"github.com/astaxie/beego/orm"     // 操作mysql的
	_ "github.com/go-sql-driver/mysql" // 驱动
)

func init() {
	// 别名连接数据库只能调用一次，再次调用会报错defaulty已经注册不能reuse
	// testmodel是操作的数据库名 mysql是使用的驱动名 default是连接数据库的一个别名
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/testmodel?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo)) // 创建一个user_info表
}

// UserInfo 这个结构体要和数据库表的名字（user_info）对应
//（数据库会自动把大写的变成小写并增加下划线）
type UserInfo struct {
	ID       int64
	Username string
	Passwd   string
}

// TestModelController .
type TestModelController struct {
	beego.Controller
}

// Get .
func (c *TestModelController) Get() {
	o := orm.NewOrm()
	user := UserInfo{Username: "hwl", Passwd: "123456"}
	id, err := o.Insert(&user) // 想要插入成功，需要事先在数据库里面创建表user_info
	if err != nil {
		logs.Println(err)
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("操作数据库id:%v", id))
}
