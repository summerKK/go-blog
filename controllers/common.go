//公共公众器

package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//初始化
func (b *BaseController) Prepare() {
	
}
