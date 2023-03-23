package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["WebSite"] = "http://192.168.10.152:8080/"
	c.Data["WebName"] = "逛逛街"
	c.Data["Email"] = "admin@xineryou.com"
	c.TplName = "index.tpl"
}
