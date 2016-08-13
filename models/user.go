package models

import (
	"github.com/astaxie/beego/orm"
)

func ReadUser(uname string) string {
	o := orm.NewOrm()
	user := User{Uname:uname}
	o.Read(&user,"Uname")
	return user.Upass
}