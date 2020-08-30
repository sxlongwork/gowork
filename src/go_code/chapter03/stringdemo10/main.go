package main

import "fmt"

func main() {
	//双引号形式
	var str1 string = "aabb\ncc"
	fmt.Println("str1 =", str1)

	//反引号`
	var str2 = `
	hello
		world
	`
	fmt.Println("str2 =", str2)

	//字符串拼接
	var str3 string = "hello " + "xianqian"
	fmt.Println("str3 =", str3)

	//当拼接的字符串很长时，可以换行，但是+号要放到行尾
	var str4 string = "hello " + "xianqian" + "hello " +
	 "xianqian" +"hello " + "xianqian" +"hello " +
	  "xianqian" +"hello " + "xianqian" +"hello " +
	   "xianqian" +"hello " + "xianqian"
	fmt.Println("str4 =", str4)

	// 基本数据类型默认值
	var a int
	var b float64
	var c bool
	var d string
	fmt.Printf("a =%d, b =%v, c =%v, d =%v", a, b, c, d)
}