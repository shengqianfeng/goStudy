package main
import (
	"fmt"
)

/**
1 go中的数据类型转换可以由小范围--》大范围，也可以大范围--》小范围
2 被转换的是变量的值，而变量本身的数据类型并未改变
3 在转换中，比如将int64转成int8，编译时不会报错，只是转换的结果会按溢出处理，最终结果变成另外一个数


*/
//演示golang中基本数据类型的转换
func main() {

	var i int32 = 100
	//希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i) //高精度-->低精度
	var n3 int64 = int64(i) //低精度->高精度

	//%v按照原值输出
	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n", i ,n1, n2, n3)
	
	//被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", i) // int32

	//在转换中，比如将 int64  转成 int8 【-128---127】 ，编译时不会报错，
	//只是转换的结果是按溢出处理，和我们希望的结果不一样
	var num1 int64 = 999999
	var num2 int8 = int8(num1) // 
	fmt.Println("num2=", num2)

}