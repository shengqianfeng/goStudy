package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//演示表单解析
func main() {
	http.HandleFunc("/", parseForm)
	http.ListenAndServe(":8080", nil)
}

const tpl = `
<html>
	<head>
		<title>hello world</title>
	</head>
	<body>
		<form method="post" action="/">
			Username: <input type="text" name="uname">
			Password: <input type="password" name="pwd">
			<button type="submit">Submit</button>
		</form>
	</body>
</html>
`

func parseForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//New allocates a new HTML template with the given name
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(w, nil)
	} else {
		//第一种方法
		//r.ParseForm()
		//第二种
		//打印出整个表单
		fmt.Println(r.Form) //返回 map[]
		//控制台打印出uname的值
		fmt.Println(r.FormValue("uname")) //admin
	}
}
