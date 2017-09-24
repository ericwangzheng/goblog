package front

import (
	"github.com/nsecgo/goblog/models"
	"strings"
)

// @router / [get]
func (c *FrontController) Index() {
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	c.Data["title"] = "nsec's Blog"
	c.Layout = "master.html"
	c.TplName = "index.html"
	articles, total := models.GetArticles(10, page)
	c.Data["articles"] = articles
	c.Data["pageNo"] = page
	c.Data["totalCount"] = total
	c.Data["tagsandcount"] = models.GetTagsAndCount()
}

// @router /articleid/:id([0-9]+) [get]
func (c *FrontController) ShowArticleById() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.GetArticleById(id)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.Layout = "master.html"
	c.TplName = "article.html"
	c.Data["title"] = article.Title
	c.Data["article"] = article
	c.Data["id"] = id
}

// @router /tag/:tag [get]
func (c *FrontController) ShowArticlesByTag() {
	tag := c.GetString(":tag")
	articles, err := models.GetArticlesByTag(tag)
	if err != nil {
		c.Redirect("/", 302)
	}
	c.Layout = "master.html"
	c.TplName = "index.html"
	c.Data["articles"] = articles
	c.Data["tag"] = tag
	c.Data["title"] = "贴有 \"" + tag + "\" 标签的文章"
}

// @router /search [get]
func (c *FrontController) Search() {
	key := strings.TrimSpace(c.GetString("key"))
	if len(key) == 0 {
		c.Redirect("/", 302)
	}
	page, err := c.GetInt("page")
	if err != nil {
		page = 1
	}
	articles, total := models.Search(key, 10, page)
	c.Layout = "master.html"
	c.TplName = "index.html"
	c.Data["articles"] = articles
	c.Data["search"] = key
	c.Data["title"] = "搜索 \"" + key + "\" 的结果"
	c.Data["pageNo"] = page
	c.Data["totalCount"] = total
}
// @router /about [get]
func (c *FrontController) About() {
	c.Layout = "master.html"
	c.TplName = "about.html"
}
