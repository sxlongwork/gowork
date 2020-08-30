package main

import "fmt"

func main() {
	//如果第一个数大于10.0，第二个数小于20.0，打印两数之和
	var n1 float64 = 20.33
	var n2 float64 = 12.3
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("n1 + n2 =", n1 + n2)
	}
	

}