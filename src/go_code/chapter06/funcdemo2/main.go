package main

import (
	"fmt"
)

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 =", n1)
}

func getSum(n1 int, n2 int) int {
	fmt.Println("test() sum =", n1 + n2)
	return n1 + n2
}
func main(){

	n1 := 10
	//调用test函数
	test(n1)	// 11
	fmt.Println("main n1 =", n1)	//10

	//调用getSum函数
	sum := getSum(10, 20)
	fmt.Println("main sum =", sum)

}