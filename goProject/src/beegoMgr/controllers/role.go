package controllers

import (
	"beegoMgr/models"
	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type RoleController struct {
	beego.Controller
}

func(c *RoleController) Get(){
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
	totalCount,err := models.GetMenuCount();
	page:=models.GetPage(pagesize,pageIndex,totalCount)

	roles,err:=models.GetAllRoleListByPage(page)
	if err != nil{
		log.Info("分页查询角色信息出错",err)
		return
	}
	log.Info("roles:",roles)

	c.Data["roles"]=roles
	//这里的结构体类型page，经测试在模板上使用既可以是指针，也可以是结构体
	c.Data["page"]=page
	c.TplName = "roleList.html"
}
func (c *RoleController) ToAddRole(){
	c.TplName = "addRole.html"
}
func (c *RoleController) ToSetPermission(){
	id,err:=c.GetInt64("id")
	if(err!=nil){
		log.Info("获取id异常，id=",id)
	}
	log.Info("开始查询菜单信息-id=",id)
	role,err:=models.GetRoleById(id)
	if(err!=nil){
		log.Info("查询角色信息异常！err=",err)
	}
	log.Info("role：",role)
	c.Data["role"]=role
	c.TplName = "setPermission.html"
}
func (c *RoleController) ToUpdateRole(){
	id,err:=c.GetInt64("id")
	if(err!=nil){
		log.Info("获取id异常，id=",id)
	}
	log.Info("开始查询菜单信息-id=",id)
	role,err:=models.GetRoleById(id)
	if(err!=nil){
		log.Info("查询角色信息异常！err=",err)
	}
	log.Info("role：",role)
	c.Data["role"]=role
	c.TplName = "updateRole.html"
}


//post方式添加角色
func (c *RoleController) AddRole() {
	rolename := c.GetString("rolename")
	if rolename==""{
		log.Info("角色名不能为空！")
		c.Data["errorMsg"]="角色名不能为空！"
		c.TplName="addRole.html"
		return
	}
	c.Data["rolename"]=rolename


	role:=models.Role{
		RoleName:rolename,
	}
	role,err:=models.AddRole(role)
	if err!=nil{
		log.Info("添加异常！")
		c.Data["errorMsg"]="添加异常！"
		c.TplName="addRole.html"
		return
	}
	c.Redirect("/role/roleList",302)
}


//post方式添加用户
func (c *RoleController) UpdateRole() {
	id,err := c.GetInt64("id")
	if err!=nil{
		log.Info("ID不能为空！")
		c.Data["errorMsg"]="Id不能为空！"
		c.TplName="updateRole.html"
		return
	}
	roleName := c.GetString("roleName")
	if roleName==""{
		log.Info("角色名不能为空！")
		c.Data["errorMsg"]="角色名不能为空！"
		c.TplName="updateRole.html"
		return
	}
	c.Data["roleName"]=roleName

	r:=models.Role{
		RoleId:id,
		RoleName:roleName,
	}
	r,err=models.UpdateRole(r)
	if err!=nil{
		log.Info("更新异常！")
		c.Data["errorMsg"]="更新异常！"
		c.TplName="updateRole.html"
		return
	}
	c.Redirect("/role/roleList",302)
}


func (c *RoleController) DeleteRole(){
	id,err:=c.GetInt64("id")
	if err != nil{
		log.Info("获取id数据错误,id=",id)
		return
	}
	err = models.DeleteRole(id)
	if(err!=nil){
		log.Info("删除失败！")
		return
	}
	//3.返回列表页面
	c.Redirect("/role/roleList",302)
}

func (c *RoleController) SetPermission(){
	menuIds:=c.GetString("menuIds")
	log.Info("menuIds:",menuIds)
	roleId,err:=c.GetInt64("roleId")
	if(err!=nil){
		log.Error("获取roleId异常：",err,roleId)
	}
	log.Info("roleId:",roleId)

	err=models.SetPermission(roleId,menuIds)
	if err!=nil{
		c.Ctx.WriteString("failed")
	}else{
		c.Ctx.WriteString("ok")
	}
}
