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

	// 插入  想要插入成功，需要事先在数据库里面创建表user_info
	// create table `user_info` (`i_d` int(11) auto_increment, `username` varchar(20) DEFAULT '', `passwd` varchar(11), primary key (i_d));
	user := UserInfo{Username: "hwl", Passwd: "123456"}
	if _, err := o.Insert(&user); err != nil {
		logs.Println(err)
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("插入数据库:%+v", user))

	// 更新
	user = UserInfo{Username: "fxp", Passwd: "123456"}
	user.ID = 1
	if _, err := o.Update(&user); err != nil {
		logs.Println(err)
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("更新数据库:%+v", user))

}
