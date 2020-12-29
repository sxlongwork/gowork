package main

import (
	"fmt"
	"sort"
)

func getChildren(c []int, s []int) int {
	sort.Ints(c)
	sort.Ints(s)
	i := 0
	j := 0
	count := 0
	for i < len(c) && j < len(s) {
		if c[i] <= s[j] {
			count++
			i++
		}
		j++
	}
	return count
}

func main() {
	var c = []int{1, 2, 3, 4}
	var s = []int{1, 2, 2, 3}
	res := getChildren(c, s)
	fmt.Println(res)
}
