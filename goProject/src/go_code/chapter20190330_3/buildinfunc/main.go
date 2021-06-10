package main
import (
	"fmt"
)

/* 
new:主要给值类型分配内存
make：主要给引用类型分配内存
*/
func main() {

	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v , num1的地址%v\n", num1, num1, &num1)

	num2 := new(int) // *int
	//num2的类型%T => *int
	//num2的值 = 地址 0xc04204c098 （这个地址是系统分配，存的值是指向另一个空间的地址）
	//num2的地址%v = 地址 0xc04206a020  (这个地址是系统分配)
	//num2指向的值 = 100
	*num2  = 100
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址%v\n num2这个指针，指向的值=%v", 
		num2, num2, &num2, *num2)
}