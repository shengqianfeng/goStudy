package main

import (
	"fmt"
)


//write Data
func writeData(intChan chan int){
	for i:=1;i<=50;i++{
		//放入数据
		intChan <- i
		fmt.Printf("writeData 写入数据%v\n",i)
	}
	//关闭
	close(intChan)
}


//readData
func readData(intChan chan int,exitChan chan bool){
	for{
		v,ok:=<-intChan
		if !ok{
			break
		}
		fmt.Printf("readData 读到数据%v\n",v)
	}
	//readData读取完数据后，即任务完成
	exitChan<-true
	close(exitChan)
}

func main()  {
	//创建两个管道
	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for{
		v,ok := <- exitChan
		fmt.Printf("exitChan 读到数据%v\n",v)
		
		//如果exitChan读不到东西则退出循环 
		//第一次输出为true，第二次为false
		// exitChan 读到数据true
		// exitChan 读到数据false
		if !ok{
			break
		}
	}
}