package main

import "fmt"

var week_days = 7

func main() {

	//还有97天放假，共几星期几天
	fmt.Printf("还有%v星期%v天放假", 97 / week_days, 97 % week_days)
}