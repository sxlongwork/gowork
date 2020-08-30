package main

import "fmt"

func main() {
	//for循环入门
	for i := 1; i <= 10; i++ {
		fmt.Println("xianqian", i)
	}
	fmt.Println("----------------------------------")

	//for循环的第二种写法
	i := 1
	for i <= 10 {
		fmt.Println("xianqian", i)
		i++
	}
	fmt.Println("----------------------------------")

	//for循环的第三种写法
	j := 1
	for {
		if j <= 10 {
			fmt.Println("xianqian", j)
		} else {
			break
		}
		j++
	}
	fmt.Println("----------------------------------")

	// //字符串遍历方式1-传统方式
	// var str string = "xianqian"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%d, %c\n", i, str[i])
	// }
	// fmt.Println("----------------------------------")

	//字符串遍历方式1-传统方式
	var str string = "xianqian好的"
	str2 := [] rune(str)	//将字符串str转换为切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d, %c\n", i, str2[i])
	}
	fmt.Println("----------------------------------")

	//字符串遍历方式2-for-range方式
	for index, value := range str {
		fmt.Printf("%d, %c \n", index, value)
	}
}