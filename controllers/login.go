package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"html/template"
	"crypto/sha256"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	uname, ok := c.GetSecureCookie("panzer", "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}
func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Data["title"] = "登录"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}
func (c *LoginController) Post() {
	uname := c.GetString("uname")
	if uname == "" {
		c.Redirect("/login?errmsg=unameisnull", 302)
	}
	upass := models.ReadUser(uname)
	if upass == "" {
		c.Redirect("/login?errmsg=upassisnull", 302)
	}
	p := []byte(c.GetString("upass"))
	pass := fmt.Sprintf("%x", sha256.Sum256(p))
	if upass == pass {
		c.SetSecureCookie("panzer", "uname", uname)
		c.Redirect("/", 302)
	} else {
		c.Redirect("/login?errmsg=upassiswrong", 302)
	}
}
func (c *LoginController) Logout() {
	c.Ctx.SetCookie("uname", "", -1)
	c.Redirect("/login", 302)
}