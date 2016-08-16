package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func init() {
	//orm.RegisterModelWithPrefix("book_", new(Book))  //带前缀的表
	orm.RegisterModel(new(Article), new(User), new(Tag))        //不带前缀的表
	orm.RegisterDataBase("default", "sqlite3", "database.sqlite")
	orm.DefaultTimeLoc = time.UTC
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}
