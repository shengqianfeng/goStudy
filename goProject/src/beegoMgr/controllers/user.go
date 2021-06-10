package controllers

import (
	"beegoMgr/models"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	log "github.com/keepeye/log4go"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	pagesize,err := c.GetInt8("pagesize")//每页条数
	if err!=nil{
		pagesize=10
	}
	log.Info("每页显示条目：",pagesize)

	pageIndex,err := c.GetInt16("pageIndex")  //当前页码
	if err != nil{
		pageIndex = 1
	}
	log.Info("当前页码：",pageIndex)
	totalCount,err := models.GetUserCount();
	page:=models.GetPage(pagesize,pageIndex,totalCount)

	users,err:=models.GetAllUserListByPage(page)
	if err != nil{
		log.Info("分页查询用户信息出错",err)
		return
	}
	log.Info("users:",users)
	//把users写进redis
	//连接redis
	conn,_:=redis.Dial("tcp",":6379",
		redis.DialKeepAlive(1*time.Second),
		redis.DialPassword("123456"),
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second))
	//关闭redis
	defer conn.Close()
	//操作redis
	//序列化操作
	//var buffer bytes.Buffer//容器
	//enc:=gob.NewEncoder(&buffer)//定义编码器
	//enc.Encode(users)//进行编码
	//conn.Do("set","users",users)//存reids储数据
	//
	////反序列化
	////连接数据库
	//buf,_ := redis.Bytes(conn.Do("get","users"))
	//dec:=gob.NewDecoder(bytes.NewReader(buf))//获取解码器
	//dec.Decode(&users)
	//log.Info(users)

	c.Data["users"]=users
	//这里的结构体类型page，经测试在模板上使用既可以是指针，也可以是结构体
	c.Data["page"]=page
	c.TplName = "userList.html"
}

func (c *UserController) ToAddUser(){
	//查询所有的角色列表
	roles:=models.GetAllRoles()
	log.Info("roles:",roles)
	c.Data["roles"]=roles
	c.TplName = "addUser.html"
}
func (c *UserController) ToUpdateUser(){
	id,err:=c.GetInt64("id")
	if(err!=nil){
		log.Info("获取id异常，id=",id)
	}
	log.Info("开始查询用户信息-id=",id)
	user,err:=models.GetUserById(id)
	if(err!=nil){
		log.Info("查询用户信息异常！err=",err)
	}
	//查询所有的角色列表
	roles:=models.GetAllRoles()
	log.Info("roles:",roles)
	c.Data["roles"]=roles
	log.Info("user：",user)
	c.Data["user"]=user
	c.TplName = "updateUser.html"
}

//post方式添加用户
func (c *UserController) AddUser() {
	username := c.GetString("username")
	nextPassword := c.GetString("nextPassword")
	password := c.GetString("password")
	sex := c.GetString("sex")
	age,_ := c.GetInt("age")
	if username==""{
		log.Info("用户名不能为空！")
		c.Data["errorMsg"]="用户名不能为空！"
		c.TplName="addUser.html"
		return
	}
	c.Data["username"]=username
	if sex==""{
		log.Info("请选择性别！")
		c.Data["errorMsg"]="请选择性别！"
		c.TplName="addUser.html"
		return
	}
	c.Data["sex"]=sex
	if age==0{
		log.Info("请选择年龄！")
		c.Data["errorMsg"]="请选择年龄！"
		c.TplName="addUser.html"
		return
	}
	c.Data["age"]=age
	if password==""{
		log.Info("请输入密码！")
		c.Data["errorMsg"]="请输入密码！"
		c.TplName="addUser.html"
		return
	}
	c.Data["password"]=password
	if nextPassword==""{
		log.Info("请再次输入密码！")
		c.Data["errorMsg"]="请再次输入密码！"
		c.TplName="addUser.html"
		return
	}
	c.Data["nextPassword"]=nextPassword
	if password!=nextPassword{
		log.Info("两次输入密码不一致！")
		c.Data["errorMsg"]="两次输入密码不一致！"
		c.TplName="addUser.html"
		return
	}
	roleId,_ := c.GetInt64("roleId")
	log.Info("roleId:",roleId)


	u:=models.User{
		Username:username,
		Password:password,
		Sex:sex,
		Age:age,
	}
	//查找Role对象
	role,err:=models.GetRoleById(roleId)
	if err!=nil{
		log.Error("角色不存在！roleId：",roleId)
		return
	}
	u.Role = role
	u,err=models.AddUser(u)
	if err!=nil{
		log.Info("添加异常！")
		c.Data["errorMsg"]="添加异常！"
		c.TplName="addUser.html"
		return
	}
	c.Redirect("/user/userList",302)
}


//post方式添加用户
func (c *UserController) UpdateUser() {
	id,err := c.GetInt64("id")
	if err!=nil{
		log.Info("ID不能为空！")
		c.Data["errorMsg"]="Id不能为空！"
		c.TplName="updateUser.html"
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")
	sex := c.GetString("sex")
	age,_ := c.GetInt("age")
	if username==""{
		log.Info("用户名不能为空！")
		c.Data["errorMsg"]="用户名不能为空！"
		c.TplName="updateUser.html"
		return
	}
	c.Data["username"]=username
	if sex==""{
		log.Info("请选择性别！")
		c.Data["errorMsg"]="请选择性别！"
		c.TplName="updateUser.html"
		return
	}
	c.Data["sex"]=sex
	if age==0{
		log.Info("请选择年龄！")
		c.Data["errorMsg"]="请选择年龄！"
		c.TplName="updateUser.html"
		return
	}
	c.Data["age"]=age
	if password==""{
		log.Info("请输入密码！")
		c.Data["errorMsg"]="请输入密码！"
		c.TplName="updateUser.html"
		return
	}
	c.Data["password"]=password

	u:=models.User{
		Id:id,
		Username:username,
		Password:password,
		Sex:sex,
		Age:age,
	}
	roleId,_ := c.GetInt64("roleId")
	log.Info("roleId:",roleId)
	//查找Role对象
	role,err:=models.GetRoleById(roleId)
	if err!=nil{
		log.Error("角色不存在！roleId：",roleId)
		return
	}
	u.Role = role
	u,err=models.UpdateUser(u)
	if err!=nil{
		log.Info("更新异常！")
		c.Data["errorMsg"]="更新异常！"
		c.TplName="updateUser.html"
		return
	}
	c.Redirect("/user/userList",302)
}

func (c *UserController) DeleteUser(){
	id,err:=c.GetInt64("id")
	if err != nil{
		log.Info("获取id数据错误,id=",id)
		return
	}
	err = models.DeleteUser(id)
	if(err!=nil){
		log.Info("删除失败！")
		return
	}
	//3.返回列表页面
	c.Redirect("/user/userList",302)
}



