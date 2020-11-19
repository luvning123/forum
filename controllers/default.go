package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetSession("Name")
	c.Data["Name"] = name
	c.TplName = "index.tpl"
}
