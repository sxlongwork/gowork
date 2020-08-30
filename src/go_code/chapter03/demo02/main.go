package main

import "fmt"

func main() {
	//使用默认值
	var i int
	fmt.Println("i =", i)

	//根据值自行判断类型
	var j = 10.1
	fmt.Println("j =", j)

	//省略var,使用:=
	name := "xianqian"
	fmt.Println("name =", name)

}