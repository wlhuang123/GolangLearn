package controllers

import "github.com/astaxie/beego"

// Arg1Controller 带参数的
type Arg1Controller struct {
	beego.Controller
}

// Get .
// http://127.0.0.1:8080/arg/?user=hwl&passwd=Hwl55555
func (c *Arg1Controller) Get() {
	user := c.GetString("user")
	c.Ctx.WriteString(user)

	c.Ctx.WriteString("\n")

	passwd := c.Input().Get("passwd")
	c.Ctx.WriteString(passwd)
}
