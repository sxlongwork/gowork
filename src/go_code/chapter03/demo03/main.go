package main

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var n3 = 300
//上面的方式也可以一次性声明，换成下面的方式
var (
	n4 = 400
	n5 = 500
	n6 = "hello"
)

func main() {
	//声明多个变量，第一种方式
	// var name, age, sex string
	// fmt.Println("name=", name, "age=", age, "sex=", sex)

	//第二种方式
	// var name, age,sex = "xianqian", "19", "girl"
	// fmt.Println("name=", name, "age=", age, "sex=", sex)

	//第三种方式
	name, age, sex := "xianqian", "19", "girl"
	fmt.Println("name=", name, "age=", age, "sex=", sex)

	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)
	fmt.Println("n4 =", n4, "n5 =", n5, "n6 =", n6)
}