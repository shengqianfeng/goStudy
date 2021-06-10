package main

import (
	"log"
	"net/http"
	"os"
)

//Http探究2
// 输出helloworld 的版本2
func main() {
	mux := http.NewServeMux()
	/*
		func (mux *ServeMux) Handle(pattern string, handler Handler)
		第二个参数是一个接口要实现ServeHTTP：
		            type Handler interface {
						ServeHTTP(ResponseWriter, *Request)
					}
	*/
	mux.Handle("/", &myHandler{})

	/*
		    底层：
			func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
				mux.Handle(pattern, HandlerFunc(handler))
			}
	*/
	mux.HandleFunc("/bye", sayBye)

	//获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//去除/static/前缀,再进行handler的转发
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	log.Println("Starting server... v2")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye this is version 2!"))
}

type myHandler struct{}

//myHandler实现一个ServeMux的方法ServeHTTP
//这里的this可以省略
func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}
