package main
import (
	"fmt"
)

func main() {
	//map的声明和注意事项 
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" 
	a["no2"] = "吴用" 
	a["no1"] = "武松" //覆盖key相同的值
	a["no3"] = "吴用"
	//map中的key是没有顺序的 
	fmt.Println(a)
}