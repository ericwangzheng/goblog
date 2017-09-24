package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//orm.RegisterModelWithPrefix("book_", new(Book))        //带前缀的表
	orm.RegisterModel(new(Article), new(User), new(Tag)) //不带前缀的表
	var dbpath string
	if beego.AppConfig.String("dbpath") == "" {
		dbpath = "sqlite.db"
	} else {
		dbpath = beego.AppConfig.String("dbpath")
	}
	err := orm.RegisterDataBase("default", "sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
	orm.Debug = false
}
