package main

import (
	"github.com/astaxie/beego"
	_ "github.com/nsecgo/goblog/routers"
)

func main() {
	//设置 logs
	//logs.SetLogger(logs.AdapterFile, `{"filename":"log/error.log","level":0,"maxdays":20}`)
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	//logs.EnableFuncCallDepth(true)
	beego.Run()
}
