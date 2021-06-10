package main
import (
	"fmt"
)

/**
交换排序算法-冒泡排序:
1 数组的冒泡排序需要进行arr.length-1轮比较
2 每一轮比较的次数逐渐递减，每一轮比较次数等于arr.length减去当前是第几轮
3 如果前一个数比后一个数大则交换位置

*/
func BubbleSort(arr *[5]int) {

	fmt.Println("排序前arr=", (*arr))
	temp := 0 //临时变量(用于做交换)


	//冒泡排序..一步一步推导出来的
	for i :=0; i < len(*arr) - 1; i++ {
		
		for j := 0; j < len(*arr) - 1 - i; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				//交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = temp
			}
		}

	}

	
	fmt.Println("排序后arr=", (*arr))

}

func main() {

	//定义数组
	arr := [5]int{24,69,80,57,13}
	//将数组传递给一个函数，完成排序

	BubbleSort(&arr)

	fmt.Println("main arr=", arr) //有序? 是有序的
}
