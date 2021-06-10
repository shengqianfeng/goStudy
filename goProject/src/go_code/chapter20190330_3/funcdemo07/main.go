package main
import (
	"fmt"
	"strings"
)

/*
闭包
   闭包就是一个函数和与其相关的引用环境组合的一个整体

*/
//累加器函数【 func (int) int】这一块是AddUpper函数的返回值
func AddUpper() func (int) int {
	var n int = 10 
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(36) // => 36 = '$'   
		fmt.Println("str=", str) // 1. str="hello$" 2. str="hello$$" 3. str="hello$$$"
		return n
	}
}
/*
解释以上函数：
1 AddUpper是一个函数，返回的数据类型时fun(int) int
2 闭包的说明
   AddUpper返回的是一个匿名函数，但是这个匿名函数引用到匿名函数外的n，因此这个匿名函数和n形成的一个整体构成闭包。
3 可以这样理解：闭包是一个类，函数是操作，n是字段，函数和n构成了闭包
4 当我们反复的调用AddUpper返回的匿名函数时,因为n是初始化1次，因此每当调用1次n就进行累加。   
5 我们要搞清楚闭包的关键：就是要分析出返回的匿名函数它引用到了哪些变量。因为函数和它引用到的变量共同构成闭包。

*/




//
// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix)  {
			return name + suffix
		}

		return name
	}
}


func makeSuffix2(suffix string, name string)  string {


	//如果 name 没有指定后缀，则加上，否则就返回原来的名字
	if !strings.HasSuffix(name, suffix)  {
		return name + suffix
	}

	return name
	
}

func main() {
	
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1))// 11 
	fmt.Println(f(2))// 13
	fmt.Println(f(3))// 16


	//测试makeSuffix 的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg") //如果使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter")) // winter.jpg
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg

	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter")) // winter.jgp
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "bird.jpg")) // bird.jpg



}