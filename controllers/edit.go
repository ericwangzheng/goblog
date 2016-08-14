package controllers

import (
	"html/template"
	"goblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Prepare() {
	uname, ok := c.GetSecureCookie("panzer", "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}

func (c *EditController) Create() {
	c.TplName = "create.html"
	c.Data["title"] = "写文章"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}

func (c *EditController) Add() {
	a := models.Article{}
	c.ParseForm(&a)
	models.Insert(&a)
	i := strconv.Itoa(a.Id)
	c.Redirect("/" + i, 302)
}
func (c *EditController) Read() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article, err := models.Show(id)
	if err == orm.ErrNoRows {
		c.Redirect("/", 302)
	}
	c.Data["article"] = article
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "create.html"
	c.Data["id"] = id
}
func (c *EditController) Update() {
	i := c.Ctx.Input.URL()
	i = i[12:]
	id, _ := strconv.Atoi(i)
	a := models.Article{}
	c.ParseForm(&a)
	a.Id = id
	models.Update(&a)
	c.Redirect("/" + i, 302)
}