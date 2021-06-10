package main
import (
	"fmt" 
)

/*
	用户键盘输入
	1 fmt.Scanln(&age)相当于把值读入变量age地址对应的空间中，当用户输入值回车后
	  函数内部会自动判断类型并转换。
	2 fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	  按照指定的格式输入，函数会判断类型并转换，转换失败则为目标类型默认值

	  两种方式推荐第一种，需要格式化则用第二种
*/
func main() {

	//要求：可以从控制台接收用户信息，【姓名，年龄，薪水, 是否通过考试 】。
	//方式1 fmt.Scanln
	//1先声明需要的变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名 ")
	// //当程序执行到 fmt.Scanln(&name),程序会停止在这里，等待用户输入，并回车 
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄 ")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水 ")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试 ")
	// fmt.Scanln(&isPass)

	// fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)

	//方式2:fmt.Scanf,可以按指定的格式输入
	fmt.Println("请输入你的姓名，年龄，薪水, 是否通过考试， 使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)


}