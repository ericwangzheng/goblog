package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)
//定义结构体，名字为表名大写，字段大写，为表的字段
type Article struct {
	Id          int       `orm:"pk;auto"` //主键，自动增长
	Title       string
	Content     string    `orm:"type(text)"`
	Author      string
	Create_time time.Time `orm:"auto_now_add;type(datatime)"`
	Update_time time.Time `orm:"auto_now;type(datatime)"`
}

func Insert(a *Article) {
	o := orm.NewOrm()
	o.Insert(a)
}

func Index() *[]*Article {
	o := orm.NewOrm()
	var articles []*Article
	o.QueryTable("article").OrderBy("-Id").All(&articles)
	return &articles
}

func Show(id int) (Article, error) {
	o := orm.NewOrm()
	article := Article{Id:id}
	err := o.Read(&article)
	return article, err
}

func Update(article *Article) error {
	o := orm.NewOrm()
	i := Article{Id:article.Id}
	err := o.Read(&i)
	if err == nil {
		_, err := o.Update(article, "Title", "Content")
		return err
	}
	return err
}
func ShowArticlesByids(ids []int) *[]*Article {
	var articles []*Article
	o := orm.NewOrm()
	o.QueryTable("Article").Filter("Id__in", ids).OrderBy("-Id").All(&articles, "Id", "Title", "Content")
	return &articles
}
func Search(key string) *[]*Article {
	var articles []*Article
	o := orm.NewOrm()
	o.QueryTable("article").Filter("Title__icontains", key).OrderBy("-Id").All(&articles, "Id", "Title", "Content")
	return &articles
}