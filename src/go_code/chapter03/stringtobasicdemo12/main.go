package main

import (
	"strconv"
	"fmt"
)
func main() {
	//string转基本数据类型
	var str string = "true"
	var flag bool

	//字符串转换位布尔类型，当然字符串必须本身为"true"或"false"
	flag, _ = strconv.ParseBool(str)
	fmt.Printf("flag数据类型是%T，flag = %v\n", flag, flag)

	var str2 string = "123445"
	var num int64
	num, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num数据类型是%T，num = %v\n", num, num)

	var str3 string = "12.234"
	var num2 float64
	num2, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("num2数据类型是%T，num2 = %v\n", num2, num2)
}