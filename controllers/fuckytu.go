package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type FuckYTUController struct {
	beego.Controller
}

func (c *FuckYTUController) List() {
	var lists []string
	bjh := c.GetString("bjh")
	bjh = strings.Replace(bjh, ".", "", -1)
	bjh = strings.Replace(bjh, "/", "", -1)
	if len(bjh) == 0 {
		files, _ := ioutil.ReadDir("static/2016")
		for _, file := range files {
			if file.IsDir() {
				lists = append(lists, file.Name())
			}
		}
	} else {
		files, _ := ioutil.ReadDir("static/2016/" + bjh)
		for _, file := range files {
			if !file.IsDir() {
				lists = append(lists, file.Name())
			}
		}
	}
	c.TplName = "list.html"
	c.Data["bjh"] = bjh
	c.Data["lists"] = lists
}