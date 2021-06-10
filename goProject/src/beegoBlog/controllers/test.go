package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	////相当于php里面的echo 函数
	c.Ctx.WriteString("<font style=\"color:red\">welcome test controller!</font>")
}
