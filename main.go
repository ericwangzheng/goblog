package main

import (
	"github.com/astaxie/beego"
	_ "github.com/nsecgo/goblog/routers"
)

func main() {
	beego.Run()
}
