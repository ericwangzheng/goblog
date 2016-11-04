package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"os"
)

type FuckYTUController struct {
	beego.Controller
}

func (c *FuckYTUController) List() {
	splat := c.GetString(":splat")
	splat = strings.Replace(splat, ".", "", -1)
	path := "static/fuckytustu/" + splat
	var lists []string
	files, ok := ioutil.ReadDir(path)
	if ok == nil {
		for _, file := range files {
			lists = append(lists, file.Name())
		}
		c.Data["show"] = false
		file, a := os.Open(path + "/" + lists[1])
		if a == nil {
			stat, b := file.Stat()
			if b == nil {
				if stat.IsDir() == false {
					c.Data["show"] = true
				}
			}
		}

	}
	c.TplName = "list.html"
	c.Data["lists"] = lists
	c.Data["path"] = path
}