package controllers

import (
	"html/template"
	"goblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/orm"
	"strings"
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

func (c *EditController) Add() {
	c.TplName = "create.html"
	c.Data["title"] = "写文章"
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Data["tagsnohas"] = models.ReadALLtag()
}

func (c *EditController) DoAdd() {
	a := models.Article{}
	c.ParseForm(&a)
	a.Author, _ = c.GetSecureCookie("panzer", "uname")
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
	c.Data["tagshas"] = a
	c.Data["tagsnohas"] = models.ReadNohas(a)
}
func (c *EditController) DoUpdate() {
	i := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(i)
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

	c.Redirect("/" + i, 302)
}
func (c *EditController)Delete() {

}