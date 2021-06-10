package utils
import (
	"fmt"
)

//将计算功能化做成一个函数，然后在需要的时候调用即可
/*
1 函数名首字母必须大写，才能被其他包访问到
2 在给一个文件打包时，该包对应一个文件夹，比如这里的utils文件夹对应的
包名就是utils，一般情况下文件夹和包名保持一致，一般为小写字母
3 当一个文件要使用其他包函数或者变量时，需要先引入对应的包
4 package指令在文件第一行，然后是import指令
5 在import包时，路径是从$GOPATH的src下开始，不用带src，编译器会自动从src下开始引入
6 在访问其他包函数变量时，语法是包名.函数名，比如这里的main.go文件中的utils.Cal
7 如果包名比较长，Go支持给包取名，注意细节，取别名后，原来的包名就不能使用了
  如果给包取了别名，则需要使用别名来访问该包的函数和变量，否则编译报错
8 在同一个包下不能有相同的函数名（同一文件夹下两个相同包名的go文件也不能有相同函数名和全局变量名），
否则包重复定义
9 如果你要编译成一个可执行文件，就需要将这个包声明为main，即package main，这个是语法规范
如果你写的是库文件，包名可以自定义


*/
var Num int = 30

func Cal(n1 float64,n2 float64,operator byte) float64 {
	/*函数的基本语法
		func 函数名(形参列表) (返回值列表){
			执行语句
			return 返回值列表
		}
	*/
	var res float64

	switch operator {
		case '+':
			res = n1+n2
		case '-':
			res = n1-n2
		case '*':
			res = n1*n2
		case '/':
			res = n1/n2
		default:
			fmt.Println("操作符号错误...") //返回值默认就是0
	}
	return res
}