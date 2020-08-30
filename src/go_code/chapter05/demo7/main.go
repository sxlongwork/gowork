package main

import "fmt"

func main() {
	//成绩判断
	var grade int
	fmt.Println("请输入golang成绩：")
	fmt.Scanln(&grade)

	if grade == 100 {
		fmt.Println("奖励一辆BMV")
	} else if grade > 80 && grade <= 99 {
		fmt.Println("奖励一台iphone7 plus")
	} else if grade >= 60 && grade <= 80 {
		fmt.Println("奖励一个iPad")
	} else {
		fmt.Println("不及格，什么都没有！")
	}

}