package main

import (
	"fmt"
)
func main() {
	//总共100,000元，

	var total float64 = 100000
	var count int =0
	for {
		if total < 0 {
			count--
			fmt.Printf("可以经过%d个路口", count)
			break
		}
		count++
		if total > 50000 {
			total = total - total * 0.05
		} else {
			total = total - 1000
		}
	}
}