package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"html/template"
)

type Default struct {
	beego.Controller
}

func (c *Default) Prepare() {
	uname, ok := c.GetSecureCookie("panzer", "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}
func (c *Default) Index() {
	c.Data["title"] = "nsec's Blog"
	c.TplName = "index.html"
	models.Index()
	c.Data["blogs"] = models.Articles
}
func (c *Default) Create() {
	c.TplName = "create.html"
	c.Data["title"] = "写文章"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
}
func (c *Default) Insert() {
	a := models.Article{}
	c.ParseForm(&a)
	models.Insert(a)
	c.Redirect("/", 302)
}
func (c *Default) Show() {
	id := c.Ctx.Input.Param(":id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.Show(i)
	if err == orm.ErrNoRows {
		c.Redirect("/", 302)
	}
	c.TplName = "article.html"
	c.Data["article"] = article
	c.Data["title"] = "文章详情"
}