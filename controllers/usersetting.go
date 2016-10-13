package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/nsecgo/goblog/models"
	"fmt"
	"crypto/sha256"
	"strings"
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
	oldpass := []byte(strings.TrimSpace(c.GetString("oldpass")))
	newpass := []byte(strings.TrimSpace(c.GetString("newpass")))
	if len(oldpass) == 0 || len(newpass) == 0 {
		c.Redirect("/?event=length-of-pass-error", 302)
	} else {
		upass := models.GetUpassByUname(uname)
		if len(upass) != 0 {
			oldsha256pass := fmt.Sprintf("%x", sha256.Sum256(oldpass))
			println(oldpass)
			println(oldsha256pass)
			println(upass)
			if oldsha256pass == upass {
				upass = fmt.Sprintf("%x", sha256.Sum256(newpass))
				models.ChangePass(uname, upass)
			}
			c.Redirect("/?event=change-pass-success", 302)
		} else {
			c.Redirect("/?event=system-error", 302)
		}
	}
}