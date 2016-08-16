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

	var cantlogin = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie("panzer", "uname")
		if ok {
			ctx.Redirect(302, "/?event=alreadylogin")
		}
	}
	beego.InsertFilter("/login",beego.BeforeExec,cantlogin)

	var auth = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie("panzer", "uname")
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}
	ns :=
		beego.NewNamespace("/admin",
			beego.NSBefore(auth),
			beego.NSRouter("/create", &controllers.EditController{}, "get:Add;post:DoAdd"),
			beego.NSRouter("/edit/:id([0-9]+)", &controllers.EditController{}, "get:Update;post:DoUpdate"),
			beego.NSRouter("/changepass",&controllers.UserSetting{}),
		)
	//注册namespace
	beego.AddNamespace(ns)
}
