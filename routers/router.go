package routers

import (
	"BeegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/coockie_test", &controllers.CoockieController{})
	beego.Router("/temple", &controllers.TempleTestController{})
	beego.Router("/api", &controllers.ApiController{})
	beego.Router("/file", &controllers.FileController{})
}
