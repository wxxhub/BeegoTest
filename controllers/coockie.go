package controllers

import (
	"github.com/astaxie/beego"
)

type CoockieController struct {
	beego.Controller
}

func (c *CoockieController) Get() {
	if c.Ctx.GetCookie("www") == "" {
		c.Ctx.WriteString("没有Cookie")
	} else {
		c.Ctx.WriteString("找到Cookie")
	}
}