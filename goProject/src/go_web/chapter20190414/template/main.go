package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	age   string
}

func main() {
	//解析模板html
	tmpl, err := template.ParseFiles("D:\\goStudy\\goProject\\src\\go_web\\chapter20190414\\template\\index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31", Title: "我的个人网站"}
	//使用p结构体渲染模板输出到控制台
	err1 := tmpl.Execute(os.Stdout, p)
	if err1 != nil {
		fmt.Println("There was an error:", err1.Error())
	}
}
