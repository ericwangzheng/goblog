package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id          int       `orm:"pk;auto"` //主键，自动增长
	Title       string
	Content     string    `orm:"type(text)"`
	User        *User     `orm:"rel(fk);on_delete(do_nothing)"`
	Create_time time.Time `orm:"auto_now_add;type(datatime)"`
	Update_time time.Time `orm:"auto_now;type(datatime)"`
	Tags        []*Tag    `orm:"rel(m2m)"`
}

func GetArticles(num, page int) ([]Article, int64) {
	offset := num * (page - 1)
	o := orm.NewOrm()
	count, _ := o.QueryTable("Article").Count()
	var articles []Article
	o.QueryTable("Article").Limit(num, offset).OrderBy("-Id").All(&articles)
	for key, article := range articles {
		prc, _ := time.LoadLocation("PRC")
		articles[key].Create_time = article.Create_time.In(prc)
		articles[key].Update_time = article.Update_time.In(prc)
	}
	return articles, count
}

func GetArticleById(id int) (*Article, error) {
	article := &Article{}
	o := orm.NewOrm()
	err := o.QueryTable("Article").Filter("Id", id).RelatedSel().One(article)
	if err == nil {
		prc, _ := time.LoadLocation("PRC")
		article.Create_time = article.Create_time.In(prc)
		article.Update_time = article.Update_time.In(prc)
		o.LoadRelated(article, "Tags")
	}
	return article, err
}
func GetArticlesByTag(tag string) ([]Article, error) {
	var articles []Article
	o := orm.NewOrm()
	_, err := o.QueryTable("Article").Filter("Tags__Tag__Name", tag).All(&articles)
	if err == nil {
		for key, article := range articles {
			prc, _ := time.LoadLocation("PRC")
			articles[key].Create_time = article.Create_time.In(prc)
			articles[key].Update_time = article.Update_time.In(prc)
		}
	}
	return articles, err
}
func Add(article *Article, tags *[]Tag) {
	o := orm.NewOrm()
	o.Read(article.User, "Uname")
	o.Insert(article)
	m2m := o.QueryM2M(article, "Tags")
	i, _ := o.QueryTable("Tag").PrepareInsert()
	for _, tag := range *tags {
		err := o.Read(&tag, "Name")
		if err != nil {
			i.Insert(&tag)
		}
		m2m.Add(tag)
	}
	i.Close() // 别忘记关闭 statement
}

func Update(article *Article, tags *[]Tag) error {
	o := orm.NewOrm()
	ok := o.QueryTable("Article").Filter("id", article.Id).Exist()
	if ok {
		_, err := o.Update(article, "Title", "Content", "Update_time")
		if err == nil {
			m2m := o.QueryM2M(article, "Tags")
			m2m.Clear()
			i, _ := o.QueryTable("Tag").PrepareInsert()
			for _, tag := range *tags {
				err := o.Read(&tag, "Name")
				if err != nil {
					i.Insert(&tag)
				}
				m2m.Add(tag)
			}
			i.Close() // 别忘记关闭 statement
		}
		return err
	} else {
		err := errors.New("non row")
		return err
	}
}
func DelArticleById(id int) {
	o := orm.NewOrm()
	o.Delete(&Article{Id: id})
	o.QueryTable("article_tags").Filter("article_id", id).Delete()
}
func Search(key string) *[]Article {
	var articles []Article
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Title__icontains", key).OrderBy("-Id").All(&articles)
	if err == nil {
		for key, article := range articles {
			o.Read(article.User)
			prc, _ := time.LoadLocation("PRC")
			articles[key].Create_time = article.Create_time.In(prc)
			articles[key].Update_time = article.Update_time.In(prc)
		}
	}
	return &articles
}
