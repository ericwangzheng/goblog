package routers

import (
	"goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Default{},"get:Index")
	beego.Router("/create", &controllers.Default{},"get:Create")
	beego.Router("/create", &controllers.Default{},"post:Insert")
	beego.Router("/:id([0-9]+)", &controllers.Default{},"get:Show")
	beego.Router("/login",&controllers.LoginController{})
}
