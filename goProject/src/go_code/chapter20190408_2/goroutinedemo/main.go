package main
import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test () hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启了一个协程，主线程继续往下走

	//main主线程结束后，不管协程有没有执行完，都随着主线程而强制退出
	for i := 1; i <= 5; i++ {
		fmt.Println(" main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}