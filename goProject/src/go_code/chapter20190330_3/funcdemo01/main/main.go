package main
import (
	"fmt"
	"go_code/chapter20190330_3/funcdemo01/utils" 
)


func main() {

	
	fmt.Printf("Num is %v \n",utils.Num)
	//请大家完成这样一个需求:
	//输入两个数,再输入一个运算符(+,-,*,/)，得到结果.。
	//分析思路....
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'

	res := utils.Cal(n1,n2,operator)
	//四舍五入保留两位小数
	fmt.Printf("res:%.2f",res)


}