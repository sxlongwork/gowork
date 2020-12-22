package main

import "fmt"

func main() {

	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	// fmt.Println(arr)
	for _, v := range arr {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}

	var arr2 = [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println(arr2)
	var arr3 = [...][2]int{{2, 3}, {3, 4}, {3, 4}}
	fmt.Println(arr3)

	//使用双for循环遍历二维数组
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Printf("%v ", arr2[i][j])
		}
		fmt.Println()
	}

	//使用for-range遍历二维数组
	for _, v1 := range arr3 {
		for _, v2 := range v1 {
			fmt.Printf("%v ", v2)
		}
		fmt.Println()
	}
}
