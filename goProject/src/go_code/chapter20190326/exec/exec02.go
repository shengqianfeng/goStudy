package exec
import "fmt"

/*
练习：
打印出num的地址，并通过num地址的指针变量修改nun的值
*/
func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int 
	ptr = &num
	*ptr = 10 //这里修改时，会导致num的值变化
	fmt.Println("num =" , num)//10

	var a_b int = 20
	fmt.Println(a_b)

	var int int = 30
	fmt.Println(int)
}