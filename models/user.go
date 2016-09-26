package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id    int    `orm:"pk;auto"` //主键，自动增长
	Uname string `orm:"unique"`
	Upass string
	Email string `orm:"unique"`
}

func ReadUser(uname string) string {
	o := orm.NewOrm()
	user := User{Uname:uname}
	o.Read(&user, "Uname")
	return user.Upass
}
func ChangePass(uname, upass string) {
	o := orm.NewOrm()
	user := User{Uname:uname, Upass:upass}
	o.Update(&user)
}