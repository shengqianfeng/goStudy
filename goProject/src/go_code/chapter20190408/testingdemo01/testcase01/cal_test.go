package cal
import (
	"fmt"
	"testing" //引入go 的testing框架包
)

/*
testing框架：
1 将xxx_test.go的文件引入
2 对xxx.go中定义的函数进行测试，即执行xxx_test.go中的TestXxx(t *testing.T)方法
*/

//编写要给测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T) {

	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		//输出日志退出
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确...")

}

func TestHello(t *testing.T) {

	fmt.Println("TestHello被调用..")

}