package controllers

import (
	"beegoMgr/models"
	"encoding/json"
	"github.com/astaxie/beego"
	log "github.com/keepeye/log4go"
)

type MenuController struct {
	beego.Controller
}

func (c *MenuController) Get() {
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

	menus,err:=models.GetAllMenuListByPage(page)
	if err != nil{
		log.Info("分页查询菜单信息出错",err)
		return
	}
	log.Info("menus:",menus)

	c.Data["menus"]=menus
	//这里的结构体类型page，经测试在模板上使用既可以是指针，也可以是结构体
	c.Data["page"]=page
	c.TplName = "menuList.html"
}
func (c *MenuController) ToAddMenu(){
	c.TplName = "addMenu.html"
}
func (c *MenuController) ToUpdateMenu(){
	id,err:=c.GetInt64("id")
	if(err!=nil){
		log.Info("获取id异常，id=",id)
	}
	log.Info("开始查询菜单信息-id=",id)
	menu,err:=models.GetMenuById(id)
	if(err!=nil){
		log.Info("查询菜单信息异常！err=",err)
	}
	log.Info("menu：",menu)
	c.Data["menu"]=menu
	c.TplName = "updateMenu.html"
}


//post方式添加菜单
func (c *MenuController) AddMenu() {
	menuname := c.GetString("menuname")
	if menuname==""{
		log.Info("菜单名不能为空！")
		c.Data["errorMsg"]="菜单名不能为空！"
		c.TplName="addMenu.html"
		return
	}
	c.Data["menuname"]=menuname
	parentId,err := c.GetInt64("parentId")
	if err!=nil{
		log.Info("请填写父菜单Id！")
		c.Data["errorMsg"]="请填写父菜单Id！"
		c.TplName="addMenu.html"
		return
	}
	c.Data["parentId"]=parentId

	url := c.GetString("url")
	if parentId!=-1 && url==""{
		log.Info("非父菜单，请填写url路径！")
		c.Data["errorMsg"]="非父菜单，请填写url路径！"
		c.TplName="addMenu.html"
		return
	}
	c.Data["url"]=url
	sort,err := c.GetInt16("sort")
	if err!=nil {
		log.Info("请输入序号！")
		c.Data["errorMsg"]="请输入序号！"
		c.TplName="addMenu.html"
		return
	}
	c.Data["sort"]=sort

	menu:=models.Menu{
		MenuName:menuname,
		ParentId:parentId,
		Url:url,
		Sort:sort,
	}
	err=models.AddMenu(menu)
	if err!=nil{
		log.Info("添加异常！")
		c.Data["errorMsg"]="添加异常！"
		c.TplName="addMenu.html"
		return
	}
	c.Redirect("/menu/menuList",302)
}


//post方式添加菜单
func (c *MenuController) UpdateMenu() {
	id,err := c.GetInt64("id")
	if err!=nil{
		log.Info("ID不能为空！")
		c.Data["errorMsg"]="Id不能为空！"
		c.TplName="updateMenu.html"
		return
	}
	menuName := c.GetString("menuName")

	if menuName==""{
		log.Info("菜单名不能为空！")
		c.Data["errorMsg"]="菜单名不能为空！"
		c.TplName="updateMenu.html"
		return
	}
	c.Data["MenuName"]=menuName

	parentId,err := c.GetInt64("parentId")
	if err!=nil{
		log.Info("请输入父菜单Id！")
		c.Data["errorMsg"]="请输入父踩点Id！"
		c.TplName="updateMenu.html"
		return
	}
	c.Data["ParentId"]=parentId
	url := c.GetString("url")
	if parentId!=-1 && url==""{
		log.Info("请输入菜单url！")
		c.Data["errorMsg"]="请输入菜单url！"
		c.TplName="updateMenu.html"
		return
	}
	c.Data["Url"]=url
	sort,err := c.GetInt16("sort")
	if err!=nil {
		log.Info("请输入序号！")
		c.Data["errorMsg"]="请输入序号！"
		c.TplName="updateMenu.html"
		return
	}
	c.Data["sort"]=sort

	m:=models.Menu{
		Id:id,
		MenuName:menuName,
		ParentId:parentId,
		Url:url,
		Sort:sort,
	}
	err=models.UpdateMenu(m)
	if err!=nil{
		log.Info("更新异常！")
		c.Data["errorMsg"]="更新异常！"
		c.TplName="updateMenu.html"
		return
	}
	c.Redirect("/menu/menuList",302)
}

func (c *MenuController) DeleteMenu(){
	id,err:=c.GetInt64("id")
	if err != nil{
		log.Info("获取id数据错误,id=",id)
		return
	}
	err = models.DeleteMenu(id)
	if(err!=nil){
		log.Info("删除失败！")
		return
	}
	//3.返回列表页面
	c.Redirect("/menu/menuList",302)
}

/*
将菜单列表转化成json数据输出
*/
func (c *MenuController)  GetAllMenus(){
	roleId,err:=c.GetInt64("roleId")
	if err!=nil{
		log.Info("获取roleId为空")
		c.Ctx.WriteString("fail")
		return
	}
	menus,_:=models.GetAllMenus(roleId)
	log.Info("menus:",menus)
	jsonBytes, _ := json.Marshal(menus)
	c.Ctx.WriteString(string(jsonBytes))
}