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
	c.Data["title"] = article.Title
	c.Data["content"] = template.HTML(article.Content)
	c.Data["id"] = id
}