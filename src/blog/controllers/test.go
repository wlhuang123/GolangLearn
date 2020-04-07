package controllers

import (
	"github.com/astaxie/beego"
)

// TestController .
type TestController struct {
	beego.Controller
}

// Get 覆盖重写beego.Controller的方法Get
func (c *TestController) Get() {
	c.Ctx.WriteString("TestController: this is get method")
}

// Post 覆盖重写beego.Controller的方法Post
// curl -d 'id=1sdf' 127.0.0.1:8080/
func (c *TestController) Post() {
	c.Ctx.WriteString("TestController: this is post method")
}
