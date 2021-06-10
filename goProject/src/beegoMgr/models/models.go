package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	//xorm可以参考wuwen.org博客
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"log"标准库的log的log.Fatal会打印并中断程序
	log "github.com/keepeye/log4go"
)

type User struct {
	Id       int64
	Username string `orm:"size(20)"`
	Password string
	Sex string `orm:"size(2)" description:(0男 1女)`
	Age int `orm:"size(2)" description:(年龄)`
	Role *Role `orm:"rel(fk)"` //一个用户拥有一个角色
}


/************************************/
//菜单
type Menu struct {
	Id       int64
	MenuName string
	Url      string
	ParentId int64
	Sort     int16
	SMenus   []*Menu `orm:"-"`
	IsChecked   bool `orm:"-"`
	Roles []*Role `orm:"reverse(many)"` //一个菜单可以属于多个角色
}

//角色-菜单：一对多
type Role struct{
	RoleId int64 `orm:"pk;auto"`
	RoleName string  `orm:"size(20)"`
	Menus []*Menu `orm:"rel(m2m);"` //一个角色可以拥有多个菜单
}

func RegisterDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("db_url"), 30)
	orm.RegisterModel(new(User), new(Menu),new(Role))
	// 是否开启调试模式 调试模式下会打印出sql语句
	orm.Debug = true
	//开启自动建表
	orm.RunSyncdb("default", false, true)
}

func GetUserByNameAndPwd(userName, password string) (User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("username", userName).Filter("password", password).One(user)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func GetUserById(id int64) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetMenuById(id int64) (*Menu, error) {
	o := orm.NewOrm()
	menu := new(Menu)
	qs := o.QueryTable("menu")
	err := qs.Filter("id", id).One(menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
func GetRoleById(id int64) (*Role, error) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable("role")
	err := qs.Filter("role_id", id).One(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func GetAllUserListByPage(page *Page) (users []User,err error){
	o := orm.NewOrm()
	start := (int16(page.PageIndex)  -1 ) * int16(page.Pagesize)
	_,err = o.QueryTable("User").Limit(page.Pagesize,start).All(&users)
	if err!=nil{
		log.Info("分页查询用户信息出错")
		return nil,err
	}
	return users,nil
}
func GetAllRoleListByPage(page *Page) (roles []Role,err error){
	o := orm.NewOrm()
	start := (int16(page.PageIndex)  -1 ) * int16(page.Pagesize)
	_,err = o.QueryTable("Role").Limit(page.Pagesize,start).All(&roles)
	if err!=nil{
		log.Info("分页查询角色信息出错")
		return nil,err
	}
	return roles,nil
}

//查询用户总数
func GetUserCount()(count int64,err error){
	o := orm.NewOrm()
	count,err = o.QueryTable("User").Count()
	return count,err
}

func AddUser(u User)(user User,err error){
	o := orm.NewOrm()
	_,err = o.Insert(&u)
	if err != nil{
		log.Info("插入数据库错误")
		return u,err
	}
	return u, nil
}

func AddRole(r Role)(role Role,err error){
	o := orm.NewOrm()
	_,err = o.Insert(&r)
	if err != nil{
		log.Info("插入数据库错误")
		return r,err
	}
	return r, nil
}

func UpdateUser(u User)(user User,err error){
	o := orm.NewOrm()
	_,err = o.Update(&u)
	if err != nil{
		log.Info("更新数据库错误")
		return u,err
	}
	return u, nil
}


func UpdateRole(r Role)(role Role,err error){
	o := orm.NewOrm()
	_,err = o.Update(&r)
	if err != nil{
		log.Info("更新数据库错误")
		return r,err
	}
	return r, nil
}

func DeleteUser(id int64)(err error){
	o := orm.NewOrm()
	user := User{Id:id}
	err = o.Read(&user)
	if err != nil{
		log.Info("查询错误")
		return err
	}
	_,err=o.Delete(&user)
	if err != nil{
		log.Info("删除错误")
		return err
	}
	return nil
}
func DeleteMenu(id int64)(err error){
	o := orm.NewOrm()
	menu := Menu{Id:id}
	err = o.Read(&menu)
	if err != nil{
		log.Info("查询错误")
		return err
	}
	_,err=o.Delete(&menu)
	if err != nil{
		log.Info("删除错误")
		return err
	}
	return nil
}
func DeleteRole(id int64)(err error){
	o := orm.NewOrm()
	role := Role{RoleId:id}
	err = o.Read(&role)
	if err != nil{
		log.Info("查询错误")
		return err
	}
	_,err=o.Delete(&role)
	if err != nil{
		log.Info("删除错误")
		return err
	}
	return nil
}

func GetAllRoles() (roles []*Role){
	o := orm.NewOrm()
	roles = make([]*Role, 0)
	o.QueryTable("role").All(&roles)
	return roles
}

/*
查询所有菜单
*/
func GetAllMenus(roleId int64) (menus []*Menu, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in f", r)
		}
	}()

	o := orm.NewOrm()
	currentRoleMenus := make([]*Menu, 0)
	//查询当前角色拥有的菜单项
	o.Raw("select a.* from menu a LEFT JOIN role_menus b on a.id = b.menu_id where b.role_id= ?  order by a.sort asc", roleId).
		QueryRows(&currentRoleMenus)

	//查询一级父菜单
	parentMenus := make([]*Menu, 0)
	qs := o.QueryTable("menu")
	qs = qs.Filter("parent_id", -1)
	_, err = qs.OrderBy("sort").All(&parentMenus)
	for _, v := range parentMenus {
		subMenus := make([]*Menu, 0)
		qs = o.QueryTable("menu")
		qs = qs.Filter("parent_id", v.Id)
		_, err = qs.OrderBy("sort").All(&subMenus)
		for _,cmenu :=range subMenus{
			for _,menu:= range  currentRoleMenus  {
				if v.Id==menu.Id{
					v.IsChecked=true
				}
				if cmenu.Id==menu.Id{
					cmenu.IsChecked=true
				}
			}
		}
		v.SMenus = subMenus
	}
	return parentMenus, err
}

//根据当前用户的角色来查询角色对应的菜单
//参考：https://studygolang.com/articles/17275?fr=sidebar
func GetCurrentUserMenus(user User)(menus []*Menu, err error){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in f", r)
		}
	}()
	o := orm.NewOrm()
	log.Info("user.Role.RoleId:",user.Role.RoleId)
	parentMenus := make([]*Menu, 0)
	o.Raw("select a.* from menu a LEFT JOIN role_menus b on a.id = b.menu_id where b.role_id= ? and  a.parent_id=-1 order by a.sort asc", user.Role.RoleId).
		QueryRows(&parentMenus)
	childMenus := make([]*Menu, 0)
	o.Raw("select a.* from menu a LEFT JOIN role_menus b on a.id = b.menu_id where b.role_id= ? and  a.parent_id!=-1 order by a.sort asc", user.Role.RoleId).
		QueryRows(&childMenus)
	//将menus参数按照父子节点的方式组装
	for _,pmenu :=range parentMenus{
		subMenus := make([]*Menu, 0)
		for _,cmenu :=range childMenus{
			if cmenu.ParentId==pmenu.Id{
				subMenus = append(subMenus, cmenu)
			}
		}
		pmenu.SMenus=subMenus
	}
	log.Info("parentMenus:",parentMenus)
	return parentMenus,nil


}

/************************菜单操作**************************************/
func GetAllMenuListByPage(page *Page) (menus []Menu,err error){
	o := orm.NewOrm()
	start := (int16(page.PageIndex)  -1 ) * int16(page.Pagesize)
	_,err = o.QueryTable("Menu").Limit(page.Pagesize,start).All(&menus)
	if err!=nil{
		log.Info("分页查询菜单信息出错")
		return nil,err
	}
	return menus,nil
}
func GetMenuCount()(count int64,err error){
	o := orm.NewOrm()
	count,err = o.QueryTable("Menu").Count()
	return count,err
}
func AddMenu(m Menu)(err error){
	o := orm.NewOrm()
	_,err = o.Insert(&m)
	if err != nil{
		log.Info("插入数据库错误")
		return err
	}
	return nil
}

func UpdateMenu(m Menu)(err error){
	o := orm.NewOrm()
	_,err = o.Update(&m)
	if err != nil{
		log.Info("更新数据库错误")
		return err
	}
	return nil
}

func SetPermission(roleId int64,menuIds string) (err error){
	o:=orm.NewOrm()
	//1 获取操作对象
	role:=Role{RoleId:roleId}
	//2 获取role的多对多操作对象
	m2m:=o.QueryM2M(&role,"Menus")
	//获取需要插入的对象
	menuIdsSplit := strings.Split(menuIds, ",")
	for _,menuId := range menuIdsSplit{
		mId,_:=strconv.ParseInt(menuId, 10, 64)
		menu:=&Menu{Id:mId}
		o.Read(menu)
		_,err=m2m.Add(menu)
		if err!=nil{
			log.Error("权限设置异常！",err)
			return nil
		}
	}
	return nil
}