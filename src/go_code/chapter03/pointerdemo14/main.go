package main

import "fmt"


func main() {
	var num int = 20
	fmt.Printf("num=%v\n", num)
	fmt.Printf("num的地址=%v\n", &num)

	var ptr *int = &num
	//通过指针修改num变量的值
	*ptr = 30
	fmt.Printf("num=%v\n", num)
}