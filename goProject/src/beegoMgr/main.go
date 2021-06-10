package main

import (
	"beegoMgr/models"
	_ "beegoMgr/routers"

	"github.com/astaxie/beego"
	"github.com/keepeye/log4go"
)

func init() {
	models.RegisterDB()
}

//beego框架的入口文件
func main() {
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//设置session过期时间为半个小时
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 1800
	log4go.LoadConfiguration("./conf/log4go.xml")
	//视图函数的映射，必须写在run之前
	beego.AddFuncMap("ShowPrePage",models.ShowPrePage)
	beego.AddFuncMap("ShowNextPage",models.ShowNextPage)
	beego.Run()
}
