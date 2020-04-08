package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/arg1", &controllers.Arg1Controller{}, "get:Get;post:Post")
	beego.Router("/arg", &controllers.ArgController{}, "get:Get;post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test/orm", &controllers.TestModelController{}, "get:Get;post:Post")
}
