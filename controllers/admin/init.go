package admin

import (
	"github.com/astaxie/beego"
)

var CookieSecret string

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Prepare() {
	uname, ok := c.GetSecureCookie(CookieSecret, "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}
