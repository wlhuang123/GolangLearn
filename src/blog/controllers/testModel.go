package controllers

import (
	"blog/models"
	"fmt"

	"github.com/astaxie/beego" // beego的web端
)

// TestModelController .
type TestModelController struct {
	beego.Controller
}

// Get .
func (c *TestModelController) Get() {
	// 插入
	user := models.UserInfo{Username: "hwl", Passwd: "123456"}
	models.AddUser(&user)
	c.Ctx.WriteString(fmt.Sprintf("插入记录到表:%+v\n", user))

	// 更新指定id=2的记录
	user = models.UserInfo{ID: 2, Username: "fxp", Passwd: "13264243343"}
	models.UpdateUserInfo(&user)
	c.Ctx.WriteString(fmt.Sprintf("更新数据库:%+v\n", user))

	// 读取指定数据库的记录
	user = models.UserInfo{ID: 2}
	models.ReadUserInfo(&user)
	c.Ctx.WriteString(fmt.Sprintf("查询表的id=2 数据:%+v\n", user))

	// 删除数据库
	user = models.UserInfo{ID: 1}
	id, _ := models.DeleteUserInfo(&user)
	c.Ctx.WriteString(fmt.Sprintf("删除数据库:%v\n", id))

	// 查询所有表的信息
	var users []models.UserInfo
	models.ReadAllUserInfo(&users)
	c.Ctx.WriteString(fmt.Sprintf("查询表的所有数据:%+v\n", users))
}
