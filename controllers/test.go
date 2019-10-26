package controllers

import (
	"github.com/astaxie/beego"
)

// TestController .
type TestController struct {
	beego.Controller
}

// Get .
func (h *TestController) Get() {
	h.Ctx.WriteString("Hello World!")
}
