package main

import (
	"fmt"
)


func cal(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main(){

	n1 := 20
	n2 := 10
	sum, sub := cal(n1, n2)
	fmt.Println("sum =", sum, "sub =", sub)

	//只想要和，不需要差，用_忽略
	sum, _ = cal(2, 3)
	fmt.Println("sum =", sum)
}