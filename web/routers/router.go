package routers

import (
	"web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/Test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/TestInput", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/TestLogin", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/TestModel", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/TestHttpLib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
	beego.Router("/TestContext", &controllers.TestContextController{}, "get:Get;post:Post")
}
