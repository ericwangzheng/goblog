package front

import (
	"crypto/sha256"
	"fmt"
	"github.com/nsecgo/goblog/models"
	"strings"
)

// @router /login [*]
func (c *FrontController) Login() {
	if c.Ctx.Input.Method() == "GET" {
		c.Layout = "master.html"
		c.TplName = "login.html"
		c.Data["title"] = "登录"
		c.Data["uname"] = ""
	}
	if c.Ctx.Input.Method() == "POST" {
		uname := strings.TrimSpace(c.GetString("uname"))
		upass := strings.TrimSpace(c.GetString("upass"))
		if len(uname) == 0 || len(upass) == 0 {
			c.Redirect("/login?errmsg=nameorpassisnull", 302)
		}
		sha256upass := fmt.Sprintf("%x", sha256.Sum256([]byte(upass)))
		upass = models.GetUpassByUname(uname)
		if upass == sha256upass {
			c.SetSecureCookie(CookieSecret, "uname", uname)
			c.Redirect("/", 302)
		} else {
			c.Redirect("/login?errmsg=passiswrong", 302)
		}
	}
}

// @router /logout [get]
func (c *FrontController) Logout() {
	c.Ctx.SetCookie("uname", "", -1)
	c.Redirect("/login", 302)
}
