package main

import "fmt"

type animal interface {
	speak() string
	eat() string
}

type cat struct {
	name string
}

func (c *cat) speak() string {
	return c.name + "....喵喵喵..."
}
func main() {
	var c cat
	c.name = "tom"
	fmt.Println(c.speak())

	// interface定义语法
	// type  接口名称 interface {
	// 	method1(参数列表) 返回值列表
	// 	method2(参数列表) 返回值列表
	// 	...
	// 	methodn(参数列表) 返回值列表
	// }
}
