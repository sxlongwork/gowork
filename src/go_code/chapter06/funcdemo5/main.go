package main

import (
	"fmt"
)

//求斐波那契数中第n个数是多少
func test(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return test(n - 1) + test(n - 2)
	}
}

// f(1)=3;f(n)=2*f(n-1)+1,求第n个数是多少
func fn(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2 * fn(n -1) + 1
	}
}


func main(){
	res := test(5)
	fmt.Println("res =", res)

	res = fn(5)
	fmt.Println("res =", res)
}