package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

type UserInfoV2 struct {
	Username string
	Password string
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		//删除cookie并重定向到首页
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	// 获取表单信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	// 验证用户名及密码
	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		maxAge := 0
		if autoLogin {
			//cookie半个小时
			maxAge = 1800
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	//重定向到首页
	this.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
