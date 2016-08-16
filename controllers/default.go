package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"html/template"
	"time"
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

	tagcount_map := make(map[string]int64)
	all := models.ReadALLtag()
	for _, tag := range all {
		tagcount_map[tag.Name] = models.ReadCountByName(tag.Name)
	}
	c.Data["tagcount"] = tagcount_map
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
	rpc, _ := time.LoadLocation("PRC")
	article.Time = article.Time.In(rpc)
	c.TplName = "article.html"
	c.Data["article"] = article
	c.Data["title"] = article.Title
	c.Data["content"] = template.HTML(article.Content)
	c.Data["id"] = id
	c.Data["tag"] = models.ReadtagByid(i)
}