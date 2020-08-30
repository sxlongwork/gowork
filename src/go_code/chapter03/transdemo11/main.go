package main

import (
	"strconv"
	"fmt"
)
func main() {
	//基本数据类型转string
	var a int = 10
	var b float32 = 2.45
	var c bool = true
	var d byte = 'f'
	var e string	//定义一个空串

	//使用第一种方式转换
	e = fmt.Sprintf("%d", a)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	e = fmt.Sprintf("%f", b)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	e = fmt.Sprintf("%t", c)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	e = fmt.Sprintf("%c", d)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	//使用第二种方式转换
	var num int = 20
	var price float64 = 23.453
	var flag bool = true

	e = strconv.FormatInt(int64(num), 10)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	//'f'表示格式，5表示小数点后保留5位，64表示这个小数是float64
	e = strconv.FormatFloat(price, 'f', 5, 64)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	e = strconv.FormatBool(flag)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)

	//还有一个转换方法，只支持int,
	var num2 int = 30
	e = strconv.Itoa(num2)
	fmt.Printf("e数据类型是%T，e =%q\n", e, e)
}