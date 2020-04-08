package controllers

import (
	"fmt"

	"github.com/astaxie/beego"         // beego的web端
	"github.com/astaxie/beego/orm"     // 操作mysql的
	_ "github.com/go-sql-driver/mysql" // 驱动
)

// UserInfo .
type UserInfo struct {
	ID       int64
	UserName string
	Passwd   string
}

// TestModelController .
type TestModelController struct {
	beego.Controller
}

// Get .
func (c *TestModelController) Get() {
	// db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hwlDB?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/testmodel?charset=utf8", 30)

	orm.RegisterModel(new(UserInfo)) // 创建一个user_info表 （数据库会自动把大写的变成小写并增加下划线）

	o := orm.NewOrm()
	user := UserInfo{UserName: "张三", Passwd: "123456"}
	id, _ := o.Insert(&user)

	//c.TplName = "index.html"
	c.Ctx.WriteString(fmt.Sprintf("操作数据库id:%v", id))

}
