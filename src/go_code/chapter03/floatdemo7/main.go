package main

import "fmt"

func main() {
	var price float32 = 3.4
	fmt.Println("price=", price)

	//不同类型，可能出现精度损失
	var n1 float32 = 12.000000323
	var n2 float64 = 12.000000323
	fmt.Println("n1 =", n1, "n2 =", n2)

}