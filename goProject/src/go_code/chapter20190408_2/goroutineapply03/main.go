package main
import (
	"fmt"
	"time"
)



//向 intChan放入 1-8000个数
func putNum(intChan chan int) {

	for i := 1; i <= 80; i++ {    
		intChan<- i
	}

	//关闭intChan
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就
// 	//放入到primeChan
func primeNum(index int,intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for 循环
	// var num int
	var flag bool // 
	for {
		if index!=0{
			time.Sleep(time.Second)
		}
		num, ok := <-intChan //intChan 取不到..
		
		if !ok { 
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num % i == 0 {//说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数就放入到primeChan
			primeChan<- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入true
	exitChan<- true	

}

func main() {

	//创建一个大小为1000的int管道
	intChan := make(chan int , 1000)

	primeChan := make(chan int, 2000)//放入结果
	//标识退出的管道
	exitChan := make(chan bool, 8) // 4个



	start := time.Now().Unix()
	
	//开启一个协程，向 intChan放入 1-8000个数
	go putNum(intChan)
	//开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就
	//放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(i,intChan, primeChan, exitChan)
	}

	//这里我们主线程，进行处理
	//开启一个协程去检查是否四个协程都已经任务结束，结束则关闭primeChan
	go func(){
		for i := 0; i < 4; i++ {
			//如果取不出来四个则阻塞等待，否则往下走
			v,ok:=<-exitChan
			if ok{
				fmt.Println("exitChan中存入：", v)
			}
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)

		//能走到这，说明我们从exitChan 取出了4个结果，就可以放心的关闭 prprimeChan
		close(primeChan)
	}()


	//遍历我们的 primeChan ,把结果取出
	for {
		fmt.Printf("素数查询开始=====")//程序开始就会立即执行这一句
		res, ok := <-primeChan //如果此处取不到数据就会阻塞住，等待primeChan有数据进去了
		if !ok{
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")


	
}