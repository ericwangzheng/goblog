package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nsecgo/goblog/models"
	"strings"
)

type Default struct {
	beego.Controller
}

func (c *Default) Prepare() {
	uname, ok := c.GetSecureCookie(CookieSecret, "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}
func (c *Default) Index() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	c.Data["title"] = "nsec's Blog"
	c.TplName = "index.html"
	article, total := models.GetArticles(10, page)
	c.Data["articles"] = article
	c.Data["pageNo"] = page
	if a := total % 10; a == 0 {
		c.Data["pageTotal"] = total / 10
	} else {
		c.Data["pageTotal"] = total / 10 + 1
	}
	c.Data["tagsandcount"] = models.GetTagsAndCount()
}

func (c *Default) ShowArticleById() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.GetArticleById(id)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.TplName = "article.html"
	c.Data["title"] = article.Title
	c.Data["article"] = article
	c.Data["id"] = id
}

func (c *Default) ShowArticlesByTag() {
	tag := c.GetString(":tag")
	articles, err := models.GetArticlesByTag(tag)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.TplName = "index.html"
	c.Data["articles"] = articles
	c.Data["tag"] = tag
	c.Data["title"] = "打有 \"" + tag + "\" 标签的文章"
}
func (c *Default)Search() {
	key := strings.TrimSpace(c.GetString("key"))
	if len(key) == 0 {
		c.Redirect("/", 302)
	}
	articles := models.Search(key)
	c.TplName = "index.html"
	c.Data["articles"] = articles
	c.Data["search"] = key
	c.Data["title"] = "搜索 \"" + key + "\" 的结果"
}