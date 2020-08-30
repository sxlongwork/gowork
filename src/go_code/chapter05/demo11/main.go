package main

import "fmt"

func main() {
	var age byte
	// switch后可以没有表达式，类似于if-else结构
	switch {
	case age == 10:
		fmt.Println("age==10")
	case age == 20:
		fmt.Println("age==20")
	default:
		fmt.Println("xianqian")
	}

	//case也可以对范围进行判断
	var scole byte =80
	switch {
	case scole > 80:
		fmt.Println("成绩优秀")
	case scole >= 60 && scole <= 80:
		fmt.Println("成绩良好")
	default:
		fmt.Println("不及格")
	}

	//switch后可以声明/定义一个变量
	switch  grade := 65;{
	case grade > 80:
		fmt.Println("成绩优秀~")
	case grade >= 60 && grade <= 80:
		fmt.Println("成绩良好~")
	default:
		fmt.Println("不及格~")
	}

	//在case语句块后面增加fallthrough，可以继续执行下一个case
	switch num:= 10; {
	case num == 10:
		fmt.Println("num等于10")
		fallthrough
	case num == 20:
		fmt.Println("num等于20")
	default:
		fmt.Println("xianqian")
	}

	//switch语句还可以被用于type-switch来判断某个interface变量中
	// 实际指向的变量类型
	var x interface{}
	var y = 10.2
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x是%T类型", i)
	case int64:
		fmt.Println("x是int类型")
	case float64:
		fmt.Println("x是float64类型")
	case bool:
		fmt.Println("x是bool类型")
	case string:
		fmt.Println("x是string类型")
	default:
		fmt.Println("x是未知类型")
	}
}