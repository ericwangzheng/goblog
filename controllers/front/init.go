package front

import (
	"github.com/astaxie/beego"
)

var CookieSecret string

type FrontController struct {
	beego.Controller
}

func (c *FrontController) Prepare() {
	uname, ok := c.GetSecureCookie(CookieSecret, "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}
