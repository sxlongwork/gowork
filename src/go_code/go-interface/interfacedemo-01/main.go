package main

import (
	"fmt"
	"strconv"
)

type animal interface {
	speak() string
	// eat() string
}

type cat struct {
	name string
}

type A interface{
	test()
	test2()
}

type B struct{
	Name string
}

type sliceInt []int

func (si sliceInt) speak() string {
	return strconv.Itoa(len(si)) + "个元素"
}

func (b B) test() {
	fmt.Println("---test---", b.Name)
}

func (b B) test2(){
	fmt.Println("---test2---", b.Name)
}

func (c cat) speak() string {
	return c.name + "....喵喵喵..."
}
func main() {
	var c = cat{}
	c.name = "tom"

	var a animal = c	// cat实现了animal的所有接口，所以可以将cat实例赋值给animal变量
	fmt.Println(a.speak())	//tom....喵喵喵...

	var ai A = B{"jack"}
	ai.test()

	var si sliceInt = make([]int, 0)
	si = append(si, 2)
	si = append(si, 4)

	var an animal = si	// sliceInt是自定义的切片类型，实现了animal接口，所以可以赋值给animal变量并调用speak方法
	res := an.speak()
	fmt.Println(res)

}
