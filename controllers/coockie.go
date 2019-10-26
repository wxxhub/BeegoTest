package controllers

import (
	"github.com/astaxie/beego"
)

// CoockieController .
type CoockieController struct {
	beego.Controller
}

// Get .
func (c *CoockieController) Get() {
	if c.Ctx.GetCookie("www") == "" {
		c.Ctx.WriteString("没有Cookie")
	} else {
		c.Ctx.WriteString("找到Cookie")
	}
}