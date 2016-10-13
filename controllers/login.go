package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"crypto/sha256"
	"fmt"
	"strings"
	"github.com/nsecgo/goblog/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Data["title"] = "登录"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Data["uname"] = ""
}
func (c *LoginController) Post() {
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
func (c *LoginController) Logout() {
	c.Ctx.SetCookie("uname", "", -1)
	c.Redirect("/login", 302)
}