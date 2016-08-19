package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func replace(in string) (out string) {
	out = strings.Replace(in, "<p>", "", -1)
	out = strings.Replace(out, "</p>", "", -1)
	return
}

func main() {
	beego.AddFuncMap("replace", replace)
	beego.Run()
}