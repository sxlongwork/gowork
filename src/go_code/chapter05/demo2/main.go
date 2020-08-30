package main

import "fmt"

func main() {
	var age byte
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("你的年龄大于等于18，该对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不大，这次放过你了。")
	}
}