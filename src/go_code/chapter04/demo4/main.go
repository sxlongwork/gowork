package main

import "fmt"

func main() {
	//输入姓名，年龄，薪水，是否通过考试
	//使用第一种方式输入fmt.Scanln()
	var name string
	var age byte
	var sal float64
	var isPass bool
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过考试")
	// fmt.Scanln(&isPass)
	// fmt.Printf("%v %v %v %v\n", name, age, sal, isPass)

	//使用第二种方式，
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，用空格分开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("%v %v %v %v\n", name, age, sal, isPass)
}