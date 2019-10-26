package controllers

import (
	"github.com/astaxie/beego"
)

// APIController .
type APIController struct {
	beego.Controller
}

// User .
type User struct {
	Name string  //首字母一定要大写
	Age  int
	Sex  string	
}

// Get .
func (a *APIController) Get() {
	a.Data["json"] = &User{
		Name : "www",
		Age : 18,
		Sex : "男",
	}

	a.ServeJSON()
}