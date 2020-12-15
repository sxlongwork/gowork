package main

import "fmt"

//struct定义，首字母大写这说明是公有的；小写说明是私有的
type student struct {
	id   int
	name string
	age  byte
}

//tag可以为结构体的成员添加说明或者标签，便于使用，这些说明可以通过反射得到
// 当定义的结构体

func main() {
	//声明与初始化,在struct中，无论使用的是指针的方式声明还是普通方式，访问其成员都使用"."，在访问的时候编译器会自动把 stu2.name 转为 (*stu2).name。
	// 方式一
	var s1 student
	s1.id = 1
	s1.name = "xiao"
	s1.age = 20
	fmt.Println(s1, s1.id, s1.name, s1.age) //{1 xiao 20}
	//方式二, s2.id =2和(*s2).id=2相同
	var s2 *student = &student{}
	(*s2).id = 2
	(*s2).name = "ming"
	(*s2).age = 30
	fmt.Println(*s2) //{1 xiao 20}
	//方式三，struct分配内存使用new，返回的是指针。
	var s3 *student = new(student)
	s3.id = 3
	s3.name = "qian"
	s3.age = 26
	fmt.Println(*s3) //{3 qian 26}
}
