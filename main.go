package main

import (
	_ "github.com/nsecgo/goblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}