package main

import (
	"fmt"
)

func main(){
	// 使用const定义常量,常量必须初始化，没与默认值一说
	// const s string	// 编译错误，必须初始化
	const num int = 4

	// 常量的值不能修改
	// num = 5		//  编译错误

	// 常量只能修身bool、数值类型(int，float系列)、string类型
	// const arr [3]int = [3]int{1,2,3}		//  编译错误
	fmt.Println(num)

	// 常量注意事项
	// 1、比较简洁的写法
	const (
		n int = 10
		name string = "xiaoming"
		flag bool = false
	)
	fmt.Println(n,name,flag)
	// 2、更专业的写法
	const (
		a = iota
		b
		c
	)
	fmt.Println(a,b,c)	// 0 1 2
	const (
		aa = iota
		bb
		cc, dd = iota, iota
	)
	fmt.Println(aa,bb,cc,dd)	// 0 1 2 2
}