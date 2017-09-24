package main

import (
	"github.com/astaxie/beego"
	_ "github.com/nsecgo/goblog/routers"
	"os"
	"io"
)

func main() {
	//检查配置文件
	_, err := os.Stat("conf/app.conf")
	if err != nil && os.IsNotExist(err) {
		srcFile, err := os.Open("conf/app.example.conf")
		if err != nil {
			panic(err)
		}
		dstFile, err := os.Create("conf/app.conf")
		if err != nil {
			panic(err)
		}
		io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()
	}
	//设置 logs
	//logs.SetLogger(logs.AdapterFile, `{"filename":"log/error.log","level":0,"maxdays":20}`)
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	//logs.EnableFuncCallDepth(true)
	beego.Run()
}
