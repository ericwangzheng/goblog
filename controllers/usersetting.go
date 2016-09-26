package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/nsecgo/goblog/models"
	"fmt"
	"crypto/sha256"
)

type UserSetting struct {
	beego.Controller
}

func (c *UserSetting)Get() {
	c.TplName = "usersetting.html"
	c.Data["uname"], _ = c.GetSecureCookie(CookieSecret, "uname")
	c.Data["title"] = "修改密码"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}
func (c *UserSetting)Post() {
	uname, _ := c.GetSecureCookie(CookieSecret, "uname")
	upass := models.ReadUser(uname)
	if upass != "" {
		p := []byte(c.GetString("oldpass"))
		pass := fmt.Sprintf("%x", sha256.Sum256(p))
		if pass == upass {
			p = []byte(c.GetString("newpass"))
			pass = fmt.Sprintf("%x", sha256.Sum256(p))
			models.ChangePass(uname, pass)
		}
	}
	c.Redirect("/?event=change-pass-success",302)
}