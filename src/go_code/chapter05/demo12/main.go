package main

import "fmt"

func main() {
	// var char byte
	// fmt.Println("请输入一个字符：")
	// fmt.Scanf("%c", &char)

	// switch char {
	// case 'a':
	// 	fmt.Printf("%c", char-32)
	// case 'b':
	// 	fmt.Printf("%c\n", char-32)
	// case 'c':
	// 	fmt.Printf("%c\n", char-32)
	// case 'd':
	// 	fmt.Printf("%c\n", char-32)
	// case 'e':
	// 	fmt.Printf("%c\n", char-32)
	// default:
	// 	fmt.Println("other")
	// }

	var scole byte
	fmt.Println("请输入成绩[0-100]：")
	fmt.Scanln(&scole)

	switch {
	case scole >= 60:
		fmt.Println("合格")
	case scole < 60:
		fmt.Println("不合格")
	default:
		fmt.Println("hehe")
	}

	var month byte
	fmt.Println("请输入月份[1-12]：")
	fmt.Scanln(&month)

	switch month {
	case 3,4,5:
		fmt.Println("春季")
	case 6,7,8:
		fmt.Println("夏季")
	case 9,10,11:
		fmt.Println("秋季")
	case 12,1,2:
		fmt.Println("冬季")
	default:
		fmt.Println("月份输入有误")
	}
}