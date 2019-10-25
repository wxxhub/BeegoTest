package controllers

import (
	"github.com/astaxie/beego"
)

// LoginController .
type LoginController struct {
	beego.Controller
}

// Get .
func (l *LoginController) Get() {
	l.TplName = "login.html"	//返回模板
}

// Post .
func (l *LoginController) Post() {
	user := l.GetString("user")
	pwd := l.GetString("pwd")
	l.Ctx.SetCookie(user, pwd)
	l.Ctx.WriteString("user:" + user + "; pwd" + pwd)
}