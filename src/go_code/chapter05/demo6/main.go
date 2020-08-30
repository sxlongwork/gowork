package main

import "fmt"

func main() {
	//判断一个年份是否是闰年
	var year int
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)

	if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
		fmt.Println(year,"年是闰年")
	}
}