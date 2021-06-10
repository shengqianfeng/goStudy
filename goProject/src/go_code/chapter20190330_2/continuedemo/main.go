package main
import "fmt"
func main() {

	//continue案例 
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue 
			}
			fmt.Println("j=", j) 
		}
	}

	// here:
	// for i:=0; i<2; i++ {
	// 	for j:=1; j<4; j++ {
	// 	if j==2 {
	// 		continue here //当j=2时直接跳到外层循环下一次去了
	// 	}
	// 	fmt.Println("i=",i,"j=",j)
	// }
	// }

	

		
}