package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nsecgo/goblog/models"
	"github.com/astaxie/beego/orm"
	"html/template"
	"time"
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
	c.Data["title"] = "nsec's Blog"
	c.TplName = "index.html"
	articles := models.Index()
	c.Data["blogs"] = articles

	tagcount_map := make(map[string]int64)
	all := models.ReadALLtag()
	for _, tag := range all {
		tagcount_map[tag.Name], _ = models.GetArticleIdBytag(tag.Name)
	}
	c.Data["tagcount"] = tagcount_map
}

func (c *Default) Show() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.Redirect("/", 302)
	}
	article, err := models.Show(id)
	if err == orm.ErrNoRows {
		c.Redirect("/", 302)
	}
	rpc, _ := time.LoadLocation("PRC")
	article.Create_time = article.Create_time.In(rpc)
	c.TplName = "article.html"
	c.Data["article"] = article
	c.Data["title"] = article.Title
	c.Data["content"] = template.HTML(article.Content)
	c.Data["id"] = id
	c.Data["tag"] = models.ReadtagByid(id)
}
func (c *Default) ReadArticleByID() {
	tag := c.Ctx.Input.Param(":tag")
	_, ids := models.GetArticleIdBytag(tag)
	var idsint []int
	for _, id := range ids {
		idsint = append(idsint, id.Article_id)
	}
	articles := models.ShowArticlesByids(idsint)
	c.TplName = "index.html"
	c.Data["blogs"] = articles
	c.Data["tag"] = tag
	c.Data["title"] = "打有 \"" + tag + "\" 标签的文章"
}
func (c *Default)Search() {
	key := c.GetString("key")
	key = strings.TrimSpace(key)
	if key == "" {
		c.Redirect("/", 302)
	}
	articles := models.Search(key)
	c.TplName = "index.html"
	c.Data["blogs"] = articles
	c.Data["search"] = key
	c.Data["title"] = "搜索 \"" + key + "\" 的结果"
}