package main
import (
	"fmt"
	//引入包的时候就已经初始化了utils包的utisl.go的init方法
	"go_code/chapter20190330_3/funcinit/utils"
)

/*
init函数的细节：
1 如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是
         全局变量定义-->init函数--->main函数
2 init函数的最重要作用就是完成一些初始化工作

		 */
 var age = test()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test() int {
	fmt.Println("test()") //1
	return 90
}

//init函数,通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...") //2
}

func main() {
	fmt.Println("main()...age=", age) //3
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}