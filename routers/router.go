package routers

import (
	"go-blog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Get("/summer", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello world"))
	})

}
