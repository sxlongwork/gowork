package main

import "fmt"

func main() {
	//输入年龄，如果大于18岁，输出"你年龄大于18岁，要对自己的行为负责"
	var age byte
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if age >=18 {
		fmt.Println("你年龄大于18岁，要对自己的行为负责")
	}

	if price := 20.44; price > 20 {
		fmt.Println("oh my god!")
	}
}