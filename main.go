package main

import (
	_ "go-blog/routers"
	"github.com/astaxie/beego"
	_ "go-blog/models"
)

func main() {
	beego.Run()
}
