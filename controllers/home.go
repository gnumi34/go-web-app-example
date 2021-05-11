package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Layout = "base/base.html"
	c.TplName = "home.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Footer"] = "base/footer.html"
	c.LayoutSections["Header"] = "base/header.html"

	args1, err1 := c.GetInt("args1")
	args2, err2 := c.GetInt("args2")
	if err1 == nil && err2 == nil {
		c.Data["Result"] = args1 + args2
	}
}
