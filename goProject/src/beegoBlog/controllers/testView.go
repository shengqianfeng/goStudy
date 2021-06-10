package controllers

import (
	"github.com/astaxie/beego"
)

type TestViewController struct {
	beego.Controller
}

func (c *TestViewController) Get() {
	// var users []models.Person
	// models.ReadUserInfo(&users)

	// c.Data["Users"] = users
	// c.Data["len"] = len(users)
	// c.TplName = "test_view.tpl"
}
