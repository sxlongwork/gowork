package main

import (
	"fmt"
)
func main() {
	var num int
	var n1 int = 0
	var n2 int = 0
	for {
		fmt.Println("请输入一个整数：")
		fmt.Scanln(&num)

		// if num > 0 {
		// 	n1++
		// 	continue
		// } else if num < 0 {
		// 	n2++
		// 	continue
		// } else {
		// 	break
		// }

		if num == 0 {
			break
		}
		if num > 0 {
			n1++
			continue
		}
		n2++
	}
	fmt.Println("n1 =", n1, "n2 =", n2)
}