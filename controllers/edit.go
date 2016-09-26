package controllers

import (
	"html/template"
	"github.com/nsecgo/goblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
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
	c.Data["tagsnohas"] = models.ReadALLtag()
	c.Data["xsrf_token"] = c.XSRFToken()
}

func (c *EditController) DoAdd() {
	a := models.Article{}
	c.ParseForm(&a)
	a.Author, _ = c.GetSecureCookie(CookieSecret, "uname")
	a.Content = strings.Replace(a.Content, "<script>", "", -1)
	a.Content = strings.Replace(a.Content, "</script>", "", -1)
	models.Insert(&a)

	var tags []models.Tag
	addtag := strings.Split(c.GetString("addtag"), "|")
	for _, name := range addtag {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name, Article_id:a.Id})
		}
	}
	tag := c.GetStrings("tag")
	for _, name := range tag {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name, Article_id:a.Id})
		}
	}
	models.Addtag(tags)

	i := strconv.Itoa(a.Id)
	c.Redirect("/" + i, 302)
}
func (c *EditController) Update() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article, err := models.Show(id)
	if err == orm.ErrNoRows {
		c.Redirect("/", 302)
	}
	c.Data["article"] = article
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "create.html"
	c.Data["title"] = "正在编辑--" + article.Title
	a := models.ReadtagByid(id)
	c.Data["tagshave"] = a
	c.Data["tagsnohave"] = models.ReadNohas(a)
	c.Data["xsrf_token"] = c.XSRFToken()
}
func (c *EditController) DoUpdate() {
	id, _ := c.GetInt(":id")
	a := models.Article{}
	c.ParseForm(&a)
	a.Id = id
	a.Content = strings.Replace(a.Content, "<script>", "", -1)
	a.Content = strings.Replace(a.Content, "</script>", "", -1)
	models.Update(&a)
	models.Deletetag(id)
	var tags []models.Tag
	addtag := strings.Split(c.GetString("addtag"), "|")
	for _, name := range addtag {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name, Article_id:a.Id})
		}
	}
	tag := c.GetStrings("tag")
	for _, name := range tag {
		name = strings.TrimSpace(name)
		if name != "" {
			tags = append(tags, models.Tag{Name:name, Article_id:a.Id})
		}
	}
	models.Addtag(tags)

	c.Redirect(beego.URLFor("Default.Show", ":id", id), 302)
}
func (c *EditController)Upload() {
	f, h, err := c.GetFile("uploadimg")
	defer f.Close()
	if err != nil {
		a := fmt.Sprintf("%s", err)
		c.Ctx.WriteString("error|" + a)
	} else {
		c.SaveToFile("uploadimg", "./static/upload/img/" + h.Filename)
		c.Ctx.WriteString(c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/static/upload/img/" + h.Filename)
	}
}
func (c *EditController)Delete() {

}