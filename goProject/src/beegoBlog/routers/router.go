package routers

import (
	"beegoBlog/controllers"
	"os"

	"github.com/astaxie/beego"
)

//默认初始化方法
func init() {
	//	beego.Router("/", &controllers.MainController{})
	//只有post或者get，(多种请求方式以逗号分隔)方式请求才能访问Test方法，其他以此类推
	// beego.Router("/post", &controllers.MainController{}, "post:Post")
	// beego.Router("/test", &controllers.TestController{}, "get:Get")
	// beego.Router("/test_input", &controllers.TestInputController{}, "get:Get;post:Post")
	// beego.Router("/test_login", &controllers.LoginController{}, "get:Login;post:Post")
	// beego.Router("/test_model", &controllers.TestModelController{}, "get:Get")
	// beego.Router("/test_httplib", &controllers.TestHttpLibController{}, "get:Get")
	// beego.Router("/test_context", &controllers.TestContextController{}, "get:Get")
	//博客项目
	beego.Router("/index", &controllers.MainController{}, "get:Index")
	beego.Router("/", &controllers.HomeController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Post")
	beego.Router("/category", &controllers.CategoryController{}, "get:Get")
	beego.Router("/topic", &controllers.TopicController{}, "get:Get;post:Post")
	//除了post和get方法外其他的可以使用自动路由规则
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{}, "post:Add")
	beego.AutoRouter(&controllers.ReplyController{})

	//附件作为静态文件的处理
	//创建附件目录 会在src下beegoWeb下生成attachment目录
	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	// beego.SetStaticPath("/attachment", "attachment")
	//作为单独的控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
