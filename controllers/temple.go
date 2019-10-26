package controllers

import (
	"github.com/astaxie/beego"
)

// TempleTestController .
type TempleTestController struct {
	beego.Controller
}

// Get .
func (t *TempleTestController) Get() {
	t.Data["MESSAGE"] = "Hello, it's me."
	t.TplName = "temple.html"
}