package main

import (
	"fmt"
	_ "strings"
)
// bubble sort 冒泡排序
// 原理: 
// 1 n个数一共需要n-1轮(外循环)
// 2 每轮将最大或最小的数移动到最后，每轮需要进行(n-1-轮数)比较次数(内循环)
// 3 每轮从第一个数开始，与相邻的数比较，大的数向后移动
func bubbleSort(arr *[10]int) {
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
// 2 每轮指定一个数，将其他的数逐次与该数比较，找到最大或最小的数
// 3 将最大或最小的数与最后或最前位置的数交换
func selectSort(arr *[10]int){
	for i:= len(arr)-1; i>= 0; i--{		// 共n-1轮
		// 指定最后位置的数为最大的数
		max := i
		for j := i -1 ; j >= 0; j--{
			// 将其他位置的数与最后位置的数进行比较，找到最大的数的位置
			if arr[j] > arr[max]{
				max = j
			}
		}
		// 如果最大的数的位置不是开始指定的那个，就交换
		if max != i {
			arr[max], arr[i] = arr[i], arr[max]
		}
	}
}

func main() {
	var arr = [10]int{14, 7, 0, 13, 9, -4, 100, -12, 2, 36}
	fmt.Printf("排序前：%v\n", arr)
	// bubbleSort(&arr)
	selectSort(&arr)
	fmt.Printf("排序后: %v\n", arr)

	// var str = make([]string, 2)
	// str = append(str, "aaa")
	// str = append(str, "bbb")
	// str2 := strings.Join(str,"")
	// fmt.Println(str2)

}
