package main

import "fmt"

func main() {
	//成绩判断
	var scole float64
	fmt.Println("请输入成绩：")
	fmt.Scanln(&scole)

	if scole <=8 {
		var sex string
		fmt.Println("请输入性别(m|f)：")
		fmt.Scanln(&sex)
		if sex == "m" {
			fmt.Println("进入男子组")
		} else {
			fmt.Println("进入女子组")
		}
	} else {
		fmt.Println("淘汰")
	}

	

}