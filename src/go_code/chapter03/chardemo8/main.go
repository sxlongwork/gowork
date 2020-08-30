package main

import "fmt"

//golang的字符使用
func main() {
	var c1 byte = 'b'

	//当直接输出byte数据类型的值时，实际输出的是对应的字符的码值
	fmt.Println("c1 =", c1)

	//如果希望输出字符，则需要使用格式化输出
	fmt.Printf("c1 = %c\n", c1)

	// var c2 byte = '倩'	//overflow
	var c2 int = '倩'
	fmt.Printf("c2 = %c, c2对应的码值%d", c2, c2)
}