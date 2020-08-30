package main

import (
	"fmt"
	"go_code/chapter06/funcdemo1/utils"
)
func main(){

	var n1 float64 = 2.3
	var n2 float64 = 9.3
	var operator byte = '+'

	res := utils.Cal(n1, n2, operator)
	fmt.Printf("res = %.2f", res)
}