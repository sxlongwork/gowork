package main

import "fmt"

//顺序查找
// func find(arr *[5]string, value string) {
// 	for i, v := range *arr {
// 		if v == value {
// 			fmt.Println("找到了")
// 			break
// 		} else if i == len(*arr)-1 {
// 			fmt.Println("没找到")
// 		}
// 	}
// }

//二分查找，使用了递归法,需要数组本身是有序的
func binaryFind(arr *[6]int, value int, start int, end int) int {
	if start > end {
		fmt.Println("没找到")
		return -1
	}
	middle := (start + end) >> 1

	if value > (*arr)[middle] {
		middle = binaryFind(arr, value, middle+1, end)
	} else if value < (*arr)[middle] {
		middle = binaryFind(arr, value, start, middle-1)
	} else {
		return middle
	}
	return middle
}

func main() {
	// var arr = [5]string{"aa", "bb", "cc", "dd", "ee"}
	// findValue := ""
	// fmt.Scanf("%s", &findValue)
	// find(&arr, findValue)

	var arr1 = [6]int{2, 4, 9, 12, 34, 45}
	index := binaryFind(&arr1, 45, 0, len(arr1)-1)
	if index != -1 {
		fmt.Println(index, arr1[index])
	}
}
