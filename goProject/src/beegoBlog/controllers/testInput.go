package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (c *TestInputController) Get() {
	//id := c.GetString("id")
	//c.Ctx.WriteString("<html>" + id + "<br/>")
	////返回了所有的请求数据
	//name := c.Input().Get("name")
	//c.Ctx.WriteString(name + "</html>")
	name := c.GetSession("name")
	password := c.GetSession("password")

	if nameString, ok := name.(string); ok && nameString != "" {
		c.Ctx.WriteString("Name:" + name.(string) + " password:" + password.(string))
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input" method="post"> 
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
	}
}

func (c *TestInputController) Post() {
	//这种方式适合拿到xml包体
	fmt.Println("===========requestBody is:", c.Ctx.Input.RequestBody)
	u := User{}
	fmt.Println("1 u is:", u)  //1 u is: { }
	err := c.ParseForm(&u)     //这一步请求数据被填充到了u实例中
	fmt.Println(" 2 u is:", u) // 2 u is: {张三 12321}
	if err != nil {
		//process error
	}
	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}
