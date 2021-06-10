package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

//Http探究3
func main() {
	/*http.Server{}中的handler参数：
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
	需要实现ServerHttp方法
	*/
	server := &http.Server{
		Addr:        ":4000",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	//自定义路由
	mux["/"] = sayHello
	mux["/bye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//判断mux中是否有访问的路由
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, " URL is: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, " Hello this version3")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "  Bye this is version 3 ")
}
