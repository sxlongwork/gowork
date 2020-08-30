package main

import "fmt"


func main() {

	//基本数据类型
	var num int = 10
	//获取num在内存中的地址，用&
	fmt.Printf("num的地址=%v\n", &num)

	//指针变量存储的是地址
	// var ptr *int = &num
	//1.ptr是指针变量
	//2.ptr的数据类型是*int
	//3.ptr本身的值是&num
	var ptr *int = &num
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)
}