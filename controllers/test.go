package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (h *TestController) Get() {
	h.Ctx.WriteString("Hello World!")
}