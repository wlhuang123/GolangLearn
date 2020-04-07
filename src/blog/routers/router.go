package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
}
