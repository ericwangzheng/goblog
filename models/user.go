package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int        `orm:"pk;auto"` //主键，自动增长
	Uname   string     `orm:"unique"`
	Upass   string
	Email   string     `orm:"unique"`
	Article []*Article `orm:"reverse(many)"`
}

func GetUpassByUname(uname string) string {
	o := orm.NewOrm()
	user := User{Uname: uname}
	o.Read(&user, "Uname")
	return user.Upass
}
func ChangePass(uname, upass string) {
	orm.NewOrm().QueryTable("User").Filter("Uname", uname).Update(orm.Params{"Upass": upass})
}
