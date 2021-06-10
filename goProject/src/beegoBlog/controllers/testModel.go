package controllers

import (
	"beegoBlog/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	log "github.com/keepeye/log4go"
)

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	//是否开启sql调试，会打印出sql语句
	//orm.Debug = true
	/*
		连接数据库
	*/
	//orm.RegisterDataBase("default",
	//"mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	//设置要操作的表Person
	//orm.RegisterModel(new(Person))

	//生成一个orm对象
	o := orm.NewOrm()
	//下面是插入
	//user := Person{Username: "zhangsan", Password: "123456"}
	//id, err := o.Insert(&user)
	//if err != nil {
	//	c.Ctx.WriteString(fmt.Sprintf("新增人员err：", err))
	//} else {
	//	c.Ctx.WriteString(fmt.Sprintf("新增人员ID：", id))
	//}

	//下面是更新
	//user := Person{Username: "zhangsan", Password: "111111"}
	//user.User_Id = 3
	//o.Update(&user)

	//下面是ORM读取
	//user1 := Person{User_Id: 3}
	//o.Read(&user1)
	//c.Ctx.WriteString(fmt.Sprintf("user info:%v", user1))

	//下面是原生读取
	//var users []Person
	//o.Raw("select * from person").QueryRows(&users)
	//c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))

	//var maps []orm.Params
	//o.Raw("select * from person").Values(&maps)
	//for _, v := range maps {
	//	//v :map[user_id:3 username:zhangsan sex:<nil> email:<nil> password:111111]
	//	c.Ctx.WriteString(fmt.Sprintf("v :%v", v))
	//}

	//采用queryBuilder方式进行读取
	var users []models.Person

	qb, _ := orm.NewQueryBuilder("mysql")

	// qb.Select("user_id", "username", "password").From("person").Where("username= 'zhangsan'").And("user_id").In("1").Limit(1)
	qb.Select("user_id", "username", "password").From("person")

	sql := qb.String()
	o.Raw(sql).QueryRows(&users)
	//打印log日志
	log.Info("users:", users)
	//user info:[{3 zhangsan 111111}]
	c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))

	//使用MVC方式调用
	// user := models.Person{Username: "wangwu", Password: "7654321"}
	// id, err := models.AddUser(&user)
	// c.Ctx.WriteString(fmt.Sprintf("call model result:%v", id, err))
}
