package controllers

import (
	"github.com/astaxie/beego"
)

type TempleTestController struct {
	beego.Controller
}

func (t *TempleTestController) Get() {
	t.Data["MESSAGE"] = "Hello, it's me."
	t.TplName = "temple.html"
}