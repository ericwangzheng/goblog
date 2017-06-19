package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"crypto/rand"
	"github.com/nsecgo/goblog/controllers/front"
	"github.com/nsecgo/goblog/controllers/admin"
)

func init() {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	front.CookieSecret = string(b)
	admin.CookieSecret = front.CookieSecret
}

func init() {
	var islogin = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie(front.CookieSecret, "uname")
		if ok {
			ctx.Redirect(302, "/?event=alreadylogin")
		}
	}
	var auth = func(ctx *context.Context) {
		_, ok := ctx.GetSecureCookie(front.CookieSecret, "uname")
		if !ok {
			ctx.Redirect(302, "/login")
		}
	}
	beego.InsertFilter("/login", beego.BeforeRouter, islogin)

	beego.Include(&front.FrontController{})
	ns :=
		beego.NewNamespace("/admin",
			beego.NSBefore(auth),
			beego.NSInclude(&admin.AdminController{}),
		)
	//注册namespace
	beego.AddNamespace(ns)
}
