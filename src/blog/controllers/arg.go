package controllers

import "github.com/astaxie/beego"

// ArgController 带参数的
type ArgController struct {
	beego.Controller
}

// Get .
// http://127.0.0.1:8080/arg/?user=hwl&passwd=Hwl55555
func (c *ArgController) Get() {
	user := c.GetString("user")
	c.Ctx.WriteString(user)

	c.Ctx.WriteString("\n")

	passwd := c.Input().Get("passwd")
	c.Ctx.WriteString(passwd)
}
