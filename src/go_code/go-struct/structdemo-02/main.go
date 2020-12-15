package main

import "fmt"

type member struct {
	age   byte
	name  string
	score float64
}

//结构体作为函数参数
func test(m member) {
	m.age = 10
	fmt.Println("test中的m=", m)
}

//结构体指针作为函数参数
func test2(m *member) {
	m.name = "xiaoming"
	fmt.Println("test2中的m=", *m)
}

func main() {
	//默认初始值
	var m1 = member{}
	fmt.Printf("age=%v,name=%v,score=%v\n", m1.age, m1.name, m1.score)
	m1.name = "hello"
	// struct作为函数参数时，函数内对struct的修改不会影响原有的struct，说明struct是值传递
	test(m1)
	fmt.Println("main中调用test后的m=", m1)
	// 要修改原有的struct的值，可以传递指针
	test2(&m1)
	fmt.Println("main中的调用test2后的m=", m1)

	//字面量初始化
	// var m2 = member{age: 12, name: "tom", score: 22.3}	//可以换行，但是分隔符必须要
	var m2 = member{
		age:   12,
		name:  "tom",
		score: 22.3, //注意如果这一行末尾没有逗号,则}不能写在下一行
	}
	fmt.Println(m2)
}
