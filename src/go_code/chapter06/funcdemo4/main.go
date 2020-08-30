package main

import (
	"fmt"
)

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("test n =", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("test2 n =", n)
	}
}

func main(){
	test(4)		//2,2,3
	test2(4)	//2
}