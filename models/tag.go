package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

//定义结构体，名字为表名大写，字段大写，为表的字段
type Tag struct {
	Id         int
	Name       string
	Article_id int
}

func ReadALLtag() []*Tag {
	o := orm.NewOrm()
	var a []*Tag
	o.QueryTable("tag").Distinct().All(&a, "Name")
	return a
}
func ReadtagByid(id int) []*Tag {
	o := orm.NewOrm()
	var a []*Tag
	o.QueryTable("tag").Filter("Article_id", id).Distinct().All(&a, "Name")
	return a
}
func ReadNohas(taghas []*Tag) []*Tag {
	o := orm.NewOrm()
	var a []*Tag
	qs := o.QueryTable("tag")
	for _, name := range taghas {
		qs = qs.Exclude("Name", name.Name)
	}
	qs.Distinct().All(&a, "Name")
	return a
}
func Addtag(tags []Tag) {
	o := orm.NewOrm()
	o.InsertMulti(10, tags)
}
func Deletetag(id int) {
	o := orm.NewOrm()
	fmt.Println("aaaaaaaaaaa")
	o.QueryTable("tag").Filter("article_id",id).Delete()
	fmt.Println("aaaaaaaaaaa")
}
func GetArticleIdBytag(name string) (int64, []*Tag) {
	o := orm.NewOrm()
	var tag  []*Tag
	a, _ := o.QueryTable("tag").Filter("Name", name).Distinct().All(&tag, "Article_id")
	return a, tag
}
