package main

import (
	"fmt"
)

//常量
func main()  {
	var num int

	//0 常量使用const修饰
	//基本类型可以声明而不赋值，可以声明后赋值
	num = 10
	//1 常量必须声明时赋值不然报以下编译错误
	//const declaration cannot have type without expression missing value in const declaration
	// const tax int
	const tax int = 0
	fmt.Println(num,tax)

	//2 常量不能再次赋值cannot assign to tax
	// tax = 1

	//3 常量只可以修饰：bool 数值类型(int，float系列)，string类型

	//这种是合法的
	const b = 9/3
	fmt.Println(b)

	//这种写法不确定b的最终值所以编译不通过const initializer num / 3 is not a constant
	// const b = num/3
	

	//const写main中是局部常量 写在main外是全局常量
	//常量简洁写法1
	const (
		x=1
		y=2
	)	
	fmt.Println(x,y)

	//还有中专业写法
	const (
		//表示m赋值为0，n在m基础上加1，t在n基础上加1
		m = iota
		n
		t
	)
	fmt.Println(m,n,t)//0 1 2


	//go中没有常量名必须字母大写的规范，可以不全大写
	//仍然通过首字母的大小写来控制常量的访问范围


}