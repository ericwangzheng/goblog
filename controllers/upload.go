package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
)

type UploadController struct {
	beego.Controller
}
func (c *UploadController)Upload() {
	_, h, err := c.GetFile("uploadimg")
	if err != nil {
		a := fmt.Sprintf("%s", err)
		c.Ctx.WriteString("error|" + a)
	} else {
		c.SaveToFile("uploadimg", "./static/upload/img/" + h.Filename)
		c.Ctx.WriteString(c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/static/upload/img/" + h.Filename)
	}
}