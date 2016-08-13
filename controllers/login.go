package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"html/template"
)

type LoginController struct {
	beego.Controller
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
	if upass == c.GetString("upass") {
		c.SetSecureCookie("panzer", "uname", uname)
		c.Redirect("/", 302)
	} else {
		c.Redirect("/login?errmsg=upassiswrong", 302)
	}
}