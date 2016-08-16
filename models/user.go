package models

import "github.com/astaxie/beego/orm"

type User struct {
	Uname string `orm:"pk"`
	Upass string
	Email string
}

func ReadUser(uname string) string {
	o := orm.NewOrm()
	user := User{Uname:uname}
	o.Read(&user, "Uname")
	return user.Upass
}
func ChangePass(uname ,upass string){
	o:=orm.NewOrm()
	user:=User{Uname:uname,Upass:upass}
	o.Update(&user)
}