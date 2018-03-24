package routers

import (
	"go-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.HomeController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get:Login")

}
