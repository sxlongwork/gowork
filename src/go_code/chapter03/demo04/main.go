package main

import "fmt"

func main() {
	//数据值可以变化，数据类型不能变化
	var i int
	i = 20
	i = 30
	i = "hehe" //
	fmt.Println("i = ", i)
}