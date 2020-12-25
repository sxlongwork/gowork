package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = make(map[int]string)
	m[1] = "ab"
	m[3] = "ad"
	m[9] = "ae"
	m[6] = "ee"

	var slice = make([]int, 1)
	for k := range m {
		slice = append(slice, k)
	}

	//对key排序
	sort.Ints(slice)

	for _, v := range slice {
		fmt.Printf("%v,%v\n", v, m[v])
	}

}
