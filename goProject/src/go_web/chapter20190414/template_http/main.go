package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var myTemplate *template.Template

type Result struct {
	output string
}

/*
自定义实现了这个接口
type Writer interface {
    Write(p []byte) (n int, err error)
}
*/
func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name  string
	Title string
	Age   int
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	//fmt.Fprintf(w, "hello ")
	var arr []Person
	p := Person{Name: "Mary001", Age: 10, Title: "我的个人网站"}
	p1 := Person{Name: "Mary002", Age: 10, Title: "我的个人网站"}
	p2 := Person{Name: "Mary003", Age: 10, Title: "我的个人网站"}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "hello world")
	/**
	 func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
		Execute方法将解析好的模板应用到data上，并将输出写入wr。
		如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。
		模板可以安全的并发执行。
	 */
	 err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("template render data:", resultWriter.output)
	//myTemplate.Execute(w, p)
	//myTemplate.Execute(os.Stdout, p)
	//file, err := os.OpenFile("C:/test.log", os.O_CREATE|os.O_WRONLY, 0755)
	//if err != nil {
	//	fmt.Println("open failed err:", err)
	//	return
	//}

}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("D:\\goStudy\\goProject\\src\\go_web\\chapter20190414\\template_http\\index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
