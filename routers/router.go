package routers

import (
	"goblog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.Default{}, "get:Index")
	beego.Router("/:id([0-9]+)", &controllers.Default{}, "get:Show")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	var auth = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie("panzer", "uname")
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}
	ns :=
		beego.NewNamespace("/admin",
			beego.NSBefore(auth),
			beego.NSRouter("/create", &controllers.EditController{}, "get:Create;post:Add"),
			beego.NSRouter("/edit/:id([0-9]+)", &controllers.EditController{}, "get:Read;post:Update"),
		)
	//注册namespace
	beego.AddNamespace(ns)
}
