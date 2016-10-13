package controllers

import (
	"html/template"
	"github.com/nsecgo/goblog/models"
	"github.com/astaxie/beego"
	"strings"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Prepare() {
	uname, ok := c.GetSecureCookie(CookieSecret, "uname")
	if !ok {
		c.Data["uname"] = ""
	} else {
		c.Data["uname"] = uname
	}
}

func (c *EditController) Add() {
	c.TplName = "create.html"
	c.Data["title"] = "写文章"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Data["tagsnothave"] = models.GetAllTags()
	c.Data["xsrf_token"] = c.XSRFToken()
}

func (c *EditController) DoAdd() {
	var article models.Article
	c.ParseForm(&article)
	uname, _ := c.GetSecureCookie(CookieSecret, "uname")
	article.User = &models.User{Uname:uname}
	article.Content = strings.Replace(article.Content, "<script>", "", -1)
	article.Content = strings.Replace(article.Content, "</script>", "", -1)
	var tags []models.Tag
	newtagstoadd := strings.Split(c.GetString("newtags"), "|")
	for _, name := range newtagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name})
		}
	}
	tagstoadd := c.GetStrings("tags")
	for _, name := range tagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name})
		}
	}
	models.Add(&article, &tags)
	c.Redirect(beego.URLFor("Default.ShowArticleById", ":id", article.Id), 302)
}
func (c *EditController) Update() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.GetArticleById(id)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.TplName = "create.html"
	c.Data["article"] = article
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Data["title"] = "正在编辑--" + article.Title
	c.Data["tagsnothave"] = models.GetNonTagsByHave(article.Tags)
	c.Data["xsrf_token"] = c.XSRFToken()
}
func (c *EditController) DoUpdate() {
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
			tags = append(tags, models.Tag{Name:name})
		}
	}
	tagstoadd := c.GetStrings("tags")
	for _, name := range tagstoadd {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name})
		}
	}
	models.Update(&article, &tags)
	c.Redirect(beego.URLFor("Default.ShowArticleById", ":id", id), 302)
}
