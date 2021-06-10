package main
import (
	"fmt"
)

func main() {
	//添加
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~" 
	fmt.Println(cities)

	/*
	演示删除
		delete是一个内置函数，如果key存在，则删除key-value
	*/
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)


	/*
	演示map的查找 
	val, ok := cities["no1"]
	返回值有两个，第一个是val为key的值，第二个ok为true和false
	true： key存在
	false：key不存在
	*/
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	//如果希望一次性删除所有的key
	//方法1. 遍历所有的key,如何逐一删除 [遍历]
	//方法2. 直接make一个新的空间，旧的空间被GC回收
	cities = make(map[string]string)
	fmt.Println(cities)

}