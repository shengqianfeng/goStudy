package controllers

import (
	"beegoMgr/models"

	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

//测试Post方式
func (c *MainController) Post() {
	c.Ctx.WriteString("hello beifengwang!")
}

//测试Post方式
func (this *MainController) Login() {
	// 获取表单信息
	userName := this.Input().Get("userName")
	password := this.Input().Get("password")
	user, err := models.GetUserByNameAndPwd(userName, password)
	if err != nil {
		log.Error(err)
		this.Data["msg"]="用户名或者密码错误"
		this.TplName="index.tpl"
	} else {
		//获取用户的角色和拥有的菜单权限

		this.SetSession("user", user)
		this.Redirect("/main", 302)
	}
	return
}
func (this *MainController) Logout() {
	// 获取表单信息
	user:=this.GetSession("user")
	log.Info("user：",user)
	this.DelSession("user")
	this.TplName="index.tpl"
}

func (this *MainController) Main() {
	this.TplName = "main.html"
	isLogin := checkAccount(this)
	if !isLogin {
		this.Redirect("/", 302)
		return
	}
	this.Data["IsLogin"] = true
	user := this.GetSession("user")
	this.Data["User"] = user
	log.Info("查询菜单列表：", user)
	menus, err := models.GetCurrentUserMenus(user.(models.User))
	if err != nil {
		log.Error(err)
	}
	this.Data["Menus"] = menus
}

func checkAccount(this *MainController) bool {
	user := this.GetSession("user")
	log.Info("session中user：", user)
	if user != nil {
		return true
	}
	return false
}
