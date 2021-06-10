package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

type MainController struct {
	baseController
}

//http://127.0.0.1:8080/index?lang=zh-CN
func (this *baseController) Prepare() {
	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

func (c *MainController) Index() {
	/*
		用户设置了模板之后系统会自动调用Render函数，
		这个函数是在beego的Controller中实现的，所以无需用户自己来调用渲染!
	*/
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// c.Data["Hi"] = c.Tr("hi")
	// c.Data["Bye"] = c.Tr("bye")
	c.Data["Hi"] = "hi"
	// c.Data["Bye"] = "bye"
	c.TplName = "index.tpl"
}

//测试Post方式
func (c *MainController) Post() {
	c.Ctx.WriteString("hello beifengwang!")
}

//测试Post方式
func (c *MainController) Login() {
	// var req map[string]interface{}
	// bufReq := c.Ctx.Input.RequestBody
	// json.Unmarshal(bufReq, &req)
	// lenReq := len(bufReq)
	// if lenReq > 1024 {
	// 	fmt.Print("req:", "\n"+string(bufReq[:64]), "...", string(bufReq[lenReq-64:]), "\n")
	// } else {
	// 	fmt.Print("req:", "\n"+string(bufReq), "\n")
	// }
	// person := models.Person{}
	// person.Username = c.Ctx.Input.Query("userName")
	// person.Password = c.Ctx.Input.Query("password")

	// sessionName := c.GetSession("userName")
	// log4go.Info("sessionName:", sessionName)
	// if person.Username == sessionName {
	// 	log4go.Info("session登陆成功！")
	// 	c.Data["userName"] = person.Username
	// 	c.Data["password"] = person.Password
	// 	c.TplName = "main.tpl"
	// 	return
	// }
	// err, p := models.GetUserInfo(&person)
	// if err != nil {
	// 	log4go.Error("查询失败！", err)
	// 	c.Data["msg"] = "查询异常！请联系管理员！"
	// 	c.TplName = "index.tpl"
	// } else if p == nil {
	// 	log4go.Info("用户名和密码不正确！", p.Username, p.Password)
	// 	c.Data["msg"] = "用户名和密码不正确！"
	// 	c.TplName = "index.tpl"
	// } else {
	// 	log4go.Info("登陆成功！")
	// 	c.Data["userName"] = p.Username
	// 	c.Data["password"] = p.Password
	// 	c.SetSession("userName", p.Username)
	// 	c.SetSession("password", p.Password)
	// 	c.TplName = "main.tpl"
	// }
}
