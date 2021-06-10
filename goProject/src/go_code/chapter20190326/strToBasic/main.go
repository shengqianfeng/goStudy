package main
import (
	"fmt"
	"strconv"
)

/**
演示golang中string转成基本数据类型,使用的是strconv包的函数：
注意：
	在将String转成基本数据类型时，要确保String能够转成有效的数据，
比如我们可以把123转成整数，但是不能把Hello转成整数，如果非要转，golang会
默认转成0

参考手册：https://studygolang.org/pkgdoc
*/
func main() {

	var str string = "true"
	var b bool
	// b, _ = strconv.ParseBool(str)
	// 说明
	// 1. strconv.ParseBool(str) 函数会返回两个值 (value bool, err error)
	// 2. 因为我只想获取到 value bool ,不想获取 err 所以我使用_忽略
	b , _ = strconv.ParseBool(str)  //string转bool  
	fmt.Printf("b type %T  b=%v\n", b, b)
	
	var str2 string = "1234590"
	var n1 int64
	var n2 int
	//String转成10进制的int64
	n1, _ = strconv.ParseInt(str2, 10, 64) //String转Int64
	n2 = int(n1)
	fmt.Printf("n1 type %T  n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	//String转成float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)


	//注意：因为没办法把hello转成整数，golang将其直接转成int64的默认值0
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)//n3没有转成变成了int64类型的默认值0
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)

	var x string = "hello"
	var y bool
	y , _ = strconv.ParseBool(x)  //string转bool  
	fmt.Printf("x type %T  x=%v\n",y, y)//x type bool  x=false
}