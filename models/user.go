package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id    int `orm:"pk;auto"` //主键，自动增长
	Uname string
	Upass string
	Email string
}

func ReadUser(uname string) string {
	o := orm.NewOrm()
	user := User{Uname:uname}
	o.Read(&user, "Uname")
	return user.Upass
}