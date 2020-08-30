package main

import "fmt"


func main() {

	//除法,如果运算的数都是整数，那么运算结果去掉小数，只保留整数
	fmt.Println(10 / 3) //3

	//如果要保留小数部分，需要由浮点数参与运算
	fmt.Println(10.0 / 4)	//2.5

	//取余的用法
	// a % b = a - a / b * b
	fmt.Println("10%3=", 10 %3 )	//1
	fmt.Println("-10%3=", -10 % 3)	//-1
	fmt.Println("10%-3=", 10 % -3)	//1
	fmt.Println("-10%-3=", -10 % -3)	//-1

	//自增++，自减--
	var num int = 10
	num++
	fmt.Println("num =", num)	//11
	num--
	fmt.Println("num =", num)	//10
}