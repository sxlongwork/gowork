package main

import "fmt"

func main() {
	//如果两个数的和能被3又能被5整除，打印提示信息
	var n1 int32 = 10
	var n2 int32 = 30
	sum := n1 + n2
	if sum % 3 == 0 && sum % 5 == 0 {
		fmt.Println("n1+n2能被3又能被5整除")
	}
}