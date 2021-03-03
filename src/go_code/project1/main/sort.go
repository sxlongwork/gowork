package main

import (
	"fmt"
	"strings"
)
// bubble sort 冒泡排序
// 原理: 
// 1 n个数一共需要n-1轮(外循环)
// 2 每轮将最大或最小的数移动到最后，每轮需要进行(n-1-轮数)比较次数(内循环)
// 3 每轮从第一个数开始，与相邻的数比较，大的数向后移动
func bubbleSort(arr *[5]int) {
	// 一共n个数排序, 需要进行n-1轮
	for i := 0; i < len(*arr)-1; i++ {
		// 每一轮将最大的数移动到最后位置，共进行n-1-i次
		for j := 0; j < len(*arr)-1-i; j++ {
			// 每一次相邻的数比较，前面的数大就交换位置
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

// selectSort 选择排序，是冒泡排序的优化
// 原理：
// 1 n个数需要进行n-1轮(外循环)
// 2 每轮找到最大的数，与最后一个数交换，每轮就交换一次位置
// 3 
func selectSort(arr *[5]int){
	for i:= len(arr)-1; i>= 0; i--{
		max := i
		for j := i -1 ; j >= 0; j--{
			if arr[j] > arr[max]{
				max = j
			}
		}
		if max != i {
			arr[max], arr[i] = arr[i], arr[max]
		}
	}
}

func main() {
	var arr = [5]int{14, 7, 0, 13, 9}
	fmt.Printf("排序前：%v\n", arr)
	// bubbleSort(&arr)
	selectSort(&arr)
	fmt.Printf("排序后: %v\n", arr)

	var str = make([]string, 2)
	str = append(str, "aaa")
	str = append(str, "bbb")
	str2 := strings.Join(str,"")
	fmt.Println(str2)

}
