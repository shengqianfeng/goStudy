package main

import (
	"github.com/beego/i18n"

	//_只想调用beegoWeb/routers的init方法

	"beegoBlog/models"
	_ "beegoBlog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/keepeye/log4go"
)

func init() {
	models.RegisterDB()
}

//beego框架的入口文件
func main() {
	i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")

	beego.AddFuncMap("i18n", i18n.Tr)

	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	//开启自动建表
	orm.RunSyncdb("default", false, true)
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	log4go.LoadConfiguration("./conf/log4go.xml")
	beego.Run()
}
