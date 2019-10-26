package models

type Users struct {
	Id int		`orm:"pk;auto;column(id)`
	Name string `orm:"size(15);column(name)`
	Pwd string  `orm:"size(15);column(pwd)`
}
