package main
import (
	"fmt"
)
/*

go语言中字符的使用规则：
	1 go中时使用byte来存储单个字符的

	2 go中的字符串由字符链接组成

	3 go中的字符是utf-8编码
	英文字符1个字节，汉字三个字节

	4 unicode编码包含了ASCII码表

	5 go中字符的本质是一个整数，直接输出时是该字符对应的ascii码
	所以需要使用Printf格式化才能输出字符本身

	6 字符类型可以运算，相当于一个整数，因为都有对应的unicode值

字符本质讨论：
	1 字符存储到计算机中需要将字符对应码值（整数）找出来：
	存储：字符   对应码值   二进制 存储
	读取：二进制  码值 字符  读取
	2 go语言的编码统一成utf8，方便统一，再也没有乱码困扰

*/
//演示golang中字符类型使用
func main() {
	
	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//当我们直接输出byte值，就是输出了的对应的字符的ASCII码值
	// 'a' ==> 
	fmt.Println("c1=", c1)//97
	fmt.Println("c2=", c2)//48

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)

	//var c3 byte = '北' //overflow溢出
	var c3 int = '北' //overflow溢出
	fmt.Println("c3=", c3)  //21271
	fmt.Printf("c3=%c c3对应码值=%d\n", c3, c3) //北

	//可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode 字符
	var c4 int = 22269 // 22269 -> '国' 120->'x'
	fmt.Println("c4=", c4)  //22269
	fmt.Printf("c4=%c\n", c4)

	//字符类型是可以进行运算的，相当于一个整数,运输时是按照码值运行
	var n1 = 10 + 'a' //  10 + 97 = 107
	fmt.Println("n1=", n1)//107
	fmt.Printf("n1=%c\n", n1) //k
}