package main

import "fmt"

func main() {
	//+号的作用
	//两边是数值
	var i, j int =1, 2
	var sum int
	sum =i + j
	fmt.Println("sum =", sum)

	//两边是字符串
	n, m := "hello", "xianqian"
	result := n + m
	fmt.Println("result =", result)

}