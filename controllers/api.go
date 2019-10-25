package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

type User struct {
	Name string  //首字母一定要大写
	Age  int
	Sex  string	
}

func (a *ApiController) Get() {
	a.Data["json"] = &User{
		Name : "www",
		Age : 18,
		Sex : "男",
	}

	a.ServeJSON()
}