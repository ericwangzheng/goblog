package routers

import (
	"github.com/nsecgo/goblog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/fuckytustu/*", &controllers.FuckYTUController{},"get:List")
	beego.Router("/", &controllers.Default{}, "get:Index")
	beego.Router("/articleid/:id([0-9]+)", &controllers.Default{}, "get:ShowArticleById")
	beego.Router("/tag/:tag", &controllers.Default{}, "get:ShowArticlesByTag")
	beego.Router("/search", &controllers.Default{}, "get:Search")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	var cantlogin = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie(controllers.CookieSecret, "uname")
		if ok {
			ctx.Redirect(302, "/?event=alreadylogin")
		}
	}
	beego.InsertFilter("/login", beego.BeforeExec, cantlogin)

	var auth = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie(controllers.CookieSecret, "uname")
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}
	ns :=
		beego.NewNamespace("/admin",
			beego.NSBefore(auth),
			beego.NSRouter("/create", &controllers.EditController{}, "get:Add;post:DoAdd"),
			beego.NSRouter("/edit/:id([0-9]+)", &controllers.EditController{}, "get:Update;post:DoUpdate"),
			beego.NSRouter("/edit/upload", &controllers.UploadController{}, "post:Upload"),
			beego.NSRouter("/changepass", &controllers.UserSetting{}),
		)
	//注册namespace
	beego.AddNamespace(ns)
}
