package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Qulaxin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "qulaxin.tpl"
}

func (c *MainController) Error() {
	c.Data["Message"] = c.GetString("msg")
	c.Data["RetUrl"] = c.GetString("returl")
	c.TplName = "error.tpl"
}
