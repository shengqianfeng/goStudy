package routers

import (
	"beegoMgr/controllers"

	"github.com/astaxie/beego"
)

//默认初始化方法
func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	//只有post或者get，(多种请求方式以逗号分隔)方式请求才能访问Test方法，其他以此类推
	beego.Router("/login", &controllers.MainController{}, "post:Login")
	beego.Router("/logout", &controllers.MainController{}, "get:Logout")
	beego.Router("/main", &controllers.MainController{}, "get:Main")
	beego.Router("/user/userList", &controllers.UserController{}, "get:Get;post:AddUser")
	beego.Router("/user/toAddUser", &controllers.UserController{}, "get:ToAddUser")
	beego.Router("/user/addUser", &controllers.UserController{}, "post:AddUser")
	beego.Router("/user/toUpdateUser", &controllers.UserController{}, "get:ToUpdateUser")
	beego.Router("/user/updateUser", &controllers.UserController{}, "post:UpdateUser")
	beego.Router(" /user/deleteUser", &controllers.UserController{}, "get:DeleteUser")

	beego.Router("/redis", &controllers.RedisController{}, "get:Get")

	beego.Router("/menu/menuList", &controllers.MenuController{}, "get:Get")
	beego.Router("/menu/toAddMenu", &controllers.MenuController{}, "get:ToAddMenu")
	beego.Router("/menu/addMenu", &controllers.MenuController{}, "post:AddMenu")
	beego.Router("/menu/toUpdateMenu", &controllers.MenuController{}, "get:ToUpdateMenu")
	beego.Router("/menu/updateMenu", &controllers.MenuController{}, "post:UpdateMenu")
	beego.Router("/menu/deleteMenu", &controllers.MenuController{}, "get:DeleteMenu")
	beego.Router("/menu/getAllMenus", &controllers.MenuController{}, "get:GetAllMenus")

	beego.Router("/role/roleList", &controllers.RoleController{}, "get:Get")
	beego.Router("/role/toAddRole", &controllers.RoleController{}, "get:ToAddRole")
	beego.Router("/role/addRole", &controllers.RoleController{}, "post:AddRole")
	beego.Router("/role/toUpdateRole", &controllers.RoleController{}, "get:ToUpdateRole")
	beego.Router("/role/updateRole", &controllers.RoleController{}, "post:UpdateRole")
	beego.Router("/role/deleteRole", &controllers.RoleController{}, "get:DeleteRole")
	beego.Router("/role/toSetPermission", &controllers.RoleController{}, "get:ToSetPermission")
	beego.Router("/role/setPermission", &controllers.RoleController{}, "post:SetPermission")
}
