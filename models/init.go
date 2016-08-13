package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)
//定义结构体，名字为表名大写，字段大写，为表的字段
type Article struct {
	Id      int `orm:"pk;auto"` //主键，自动增长
	Title   string
	Content string
}
type User struct {
	Id    int `orm:"pk;auto"` //主键，自动增长
	Uname string
	Upass string
	Email string
}

var Articles []*Article
//注册模型
func init() {
	//orm.RegisterModelWithPrefix("book_", new(Book))  //带前缀的表
	orm.RegisterModel(new(Article), new(User))        //不带前缀的表
	orm.RegisterDataBase("default", "sqlite3", "database.sqlite")
	orm.DefaultTimeLoc = time.Local
	orm.RunSyncdb("default", false, true)
}
func Insert(a Article) {
	o := orm.NewOrm()
	o.Insert(&a)
}
func Index() {
	o := orm.NewOrm()
	o.QueryTable("article").All(&Articles)
}
func Show(id int) (Article , error) {
	o := orm.NewOrm()
	article := Article{Id:id}
	err := o.Read(&article)
	return article, err
}
