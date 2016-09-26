package controllers

import "github.com/astaxie/beego"

var CookieSecret = beego.AppConfig.String("cookiesecret")