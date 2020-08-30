package main

import "fmt"

func main() {
	var char byte
	fmt.Println("请输入一个字符：")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'd':
		fmt.Println("星期四")
	case 'e':
		fmt.Println("星期五")
	case 'f':
		fmt.Println("星期六")
	case 'g':
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误")
	}
}