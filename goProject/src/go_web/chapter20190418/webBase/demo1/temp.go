package main

import (
	"io"
	//标准库log包实现了简单的日志服务
	"log"
	"net/http"
)

/*
  beego是基于http服务器的高层封装
*/
//Http探究1
//beego的Helloworld入门示例
// type HomeController struct {
// 	beego.Controller
// }

// func (this *HomeController) Get() {
//因为在beego.Controller中的 Ctx  *context.Context实现了WriteString方法
// 	this.Ctx.WriteString("Hello world")
// }

func main() {
	//注册controller
	// beego.Router("/", &HomeController{})
	// beego.Run()

	/*
		1 type HandlerFunc func(ResponseWriter, *Request)

		2 HandleFunc func(pattern string, handler func(ResponseWriter, *Request))


		3 标准库的net/http包的HandleFunc方法：
		HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。
	*/
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world,this version 1")
}
