package admin

import (
	"crypto/sha256"
	"fmt"
	"github.com/nsecgo/goblog/models"
	"strings"
)

// @router /changepass [get]
func (c *AdminController) Get() {
	c.Layout = "master.html"
	c.TplName = "usersetting.html"
	c.Data["uname"], _ = c.GetSecureCookie(CookieSecret, "uname")
	c.Data["title"] = "修改密码"
}

// @router /changepass [post]
func (c *AdminController) Post() {
	uname, _ := c.GetSecureCookie(CookieSecret, "uname")
	oldpass := []byte(strings.TrimSpace(c.GetString("oldpass")))
	newpass := []byte(strings.TrimSpace(c.GetString("newpass")))
	if len(oldpass) == 0 || len(newpass) == 0 {
		c.Redirect("/?event=length-of-pass-error", 302)
	} else {
		upass := models.GetUpassByUname(uname)
		if len(upass) != 0 {
			oldsha256pass := fmt.Sprintf("%x", sha256.Sum256(oldpass))
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
