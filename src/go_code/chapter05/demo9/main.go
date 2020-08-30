package main

import "fmt"

func main() {
	//根据季节不同票价不同
	var month byte
	var age byte
	var price float64
	fmt.Println("请输入当前月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		price = 60.0
		if age > 60 {
			fmt.Println("老人1/3票价", price / 3)
		} else if age >= 18 {
			fmt.Println("成人票价", price)
		} else {
			fmt.Println("儿童半价", price / 2)
		}
	} else {
		
		if age >= 18 && age <= 60 {
			price = 40
			fmt.Println("成人票价", price)
		} else {
			price = 20
			fmt.Println("其他人票价", price)
		}
	}
}