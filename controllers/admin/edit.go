package admin

import (
	"github.com/astaxie/beego"
	"github.com/nsecgo/goblog/models"
	"strings"
)

// @router /create [get]
func (c *AdminController) Add() {
	c.Layout = "master.html"
	c.TplName = "create.html"
	c.Data["title"] = "写文章"
	c.Data["tagsnothave"] = models.GetAllTags()
}

// @router /create [post]
func (c *AdminController) DoAdd() {
	var article models.Article
	c.ParseForm(&article)
	uname, _ := c.GetSecureCookie(CookieSecret, "uname")
	article.User = &models.User{Uname: uname}
	article.Content = strings.Replace(article.Content, "<script>", "", -1)
	article.Content = strings.Replace(article.Content, "</script>", "", -1)
	var tags []models.Tag
	newtagstoadd := strings.Split(c.GetString("newtags"), "|")
	for _, name := range newtagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name: name})
		}
	}
	tagstoadd := c.GetStrings("tags")
	for _, name := range tagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name: name})
		}
	}
	models.Add(&article, &tags)
	c.Redirect(beego.URLFor("FrontController.ShowArticleById", ":id", article.Id), 302)
}

// @router /edit/:id([0-9]+) [get]
func (c *AdminController) Update() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.GetArticleById(id)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.Layout = "master.html"
	c.TplName = "create.html"
	c.Data["article"] = article
	c.Data["title"] = "正在编辑--" + article.Title
	c.Data["tagsnothave"] = models.GetNonTagsByHave(article.Tags)
}

// @router /edit/:id([0-9]+) [post]
func (c *AdminController) DoUpdate() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	var article models.Article
	c.ParseForm(&article)
	article.Id = id
	article.Content = strings.Replace(article.Content, "<script>", "", -1)
	article.Content = strings.Replace(article.Content, "</script>", "", -1)
	var tags []models.Tag
	newtagstoadd := strings.Split(c.GetString("newtags"), "|")
	for _, name := range newtagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name: name})
		}
	}
	tagstoadd := c.GetStrings("tags")
	for _, name := range tagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name: name})
		}
	}
	models.Update(&article, &tags)
	c.Redirect(beego.URLFor("FrontController.ShowArticleById", ":id", id), 302)
}

// @router /delete/:id([0-9]+) [get]
func (c *AdminController) Delete() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	models.DelArticleById(id)
	c.Redirect(beego.URLFor("FrontController.Index", ":id", id), 302)
}
