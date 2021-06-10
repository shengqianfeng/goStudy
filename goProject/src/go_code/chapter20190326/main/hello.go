//开发一个程序，输出Helloworld
/*
1 go文件的后缀是.go
2 package main 表示该文件所在包是main，go程序的每一个文件必须所属于一个包不能独立存在
3 import “fmt” 表示引入一个包即fmt，目的是引入该包后就可以fmt包的函数
4 func main {

  }
  func是一个关键字，表示后面是一个函数，
  main是函数名，是一个主函数，即程序的入口
5 go严格区分大小写
6 go程序由一条条语句构成，每条语句后不需要加分号，go会默认给每条语句加上
7 Go编译器是一行行进行编译的，一行只写一条语句，不能把多行写成一行，否则报错
8 go语言声明的变量或者import包如果没有用到，则编译不能通过
9 大括号都是成对出现的，缺一不可


*/
package main
import "fmt"
func main(){
  // 输出Hello
  fmt.Println("Hello world!")
  //使用逗号，链接过长字符串
  fmt.Println("Hello world!Hello world!",
                "Hello world!Hello world!",
                "Hello world!Hello world!")

  
}

