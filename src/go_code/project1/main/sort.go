package main

import (
	"fmt"
	"strings"
)
// bubble sort
func bubbleSort(arr *[5]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	var arr = [5]int{14, 7, 0, 13, 9}
	fmt.Printf("排序前：%v\n", arr)
	bubbleSort(&arr)
	fmt.Printf("排序后: %v\n", arr)

	var str = make([]string, 2)
	str = append(str, "aaa")
	str = append(str, "bbb")
	str2 := strings.Join(str,"")
	fmt.Println(str2)

}
