package controllers

import (
	"hwl/tool/logs"

	"github.com/astaxie/beego"
)

// User .
type User struct {
	ID       int    `form:"-"` // 不解析
	Username string `form:"username"`
	Age      string `form:"age"`
	Email    string `form:"email"`
}

// ArgController 带参数的
type ArgController struct {
	beego.Controller
}

// Get .
func (c *ArgController) Get() {
	c.TplName = "index.html"
}

// Post .
func (c *ArgController) Post() {
	var u User
	if err := c.ParseForm(&u); err != nil {
		logs.Println(err)
		return
	}

	c.Ctx.WriteString("uname:" + u.Username + " age:" + u.Age + " email:" + u.Email)

}
