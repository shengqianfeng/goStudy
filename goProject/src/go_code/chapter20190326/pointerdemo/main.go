package main
import (
	"fmt"
)

//演示golang中指针类型
/*
	1 基本数据类型，变量存的就是值，也叫值类型。变量和值是同一个地址
	2 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值。

指针的使用细节：
	1 值类型，都有对应的指针类型，形式为 *数据类型，比如int对应的指针就是*int，float32对应
	的指针类型就是*float32，一次类推
	2 值类型包括，int系列，float系列，bool，string、数组和结构体struct


常见的值类型和引用类型
	1 值类型：基本数据类型int系列，float系列，bool，String、数组和结构体struct
	2 引用类型：指针、slice切片、map、管道chann、interface等都是引用类型

值类型和引用类型的使用特点：
	1 值类型：变量直接存储，内存通常在栈中分配
	2 引用类型：变量存储的是一个地址，这个地址对应的空间才是真正存储数据值，内存通常在堆上
	分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC回收。

	*/
func main() {

	//基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么,&i
	fmt.Println("i的地址=", &i)
	
	//下面的 var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身的值&i
	var ptr *int = &i 
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址=%v", &ptr) 
	fmt.Printf("ptr 指向的值=%v", *ptr)//10

}