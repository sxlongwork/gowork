package main

import "fmt"

func main() {
	//交换a,b的值，不允许使用中间变量
	var a int = 10
	var b int = 20
	fmt.Println("a =", a, "b =", b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a =", a, "b =", b)

}