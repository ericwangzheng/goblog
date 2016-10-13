package models

import (
	"github.com/astaxie/beego/orm"
)

//定义结构体，名字为表名大写，字段大写，为表的字段
type Tag struct {
	Id       int        `orm:"pk;auto"` //主键，自动增长
	Name     string     `orm:"unique"`
	Articles []*Article `orm:"reverse(many)"`
}

func GetAllTags() []Tag {
	o := orm.NewOrm()
	var tags []Tag
	o.QueryTable("Tag").All(&tags, "Name")
	return tags
}
func GetNonTagsByHave(tags []*Tag) *[]Tag {
	var have []string
	for _, tag := range tags {
		have = append(have, tag.Name)
	}
	var nontags []Tag
	orm.NewOrm().QueryTable("Tag").Exclude("Name__in", have).All(&nontags)
	return &nontags
}
func GetTagsAndCount() map[string]int64 {
	tagsandcount_map := make(map[string]int64)
	o := orm.NewOrm()
	var tags []Tag
	o.QueryTable("Tag").All(&tags)
	for _, tag := range tags {
		tagsandcount_map[tag.Name], _ = o.QueryM2M(&tag, "Articles").Count()
	}
	return tagsandcount_map
}
