package main

import (
	"flag"
	"fmt"
)

//定义命令行参数对应的变量，这两个变量名都是指针类型
var (
	name = flag.String("name", "tom", "input your name")
	age  = flag.Int("age", 21, "input your age")
)

//也可以定义值类型的变量，在init函数中完成初始化
var sex string

func init() {
	flag.StringVar(&sex, "sex", "male", "input your sex")
}

func main() {
	//把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	//可以通过 flag.Args() 和 flag.NArg() 函数获取未能解析的命令行参数。

	fmt.Printf("name=%v, age=%v, sex=%v", *name, *age, sex)
	//命令行传参的格式：
	// -isbool		(一个 - 符号，布尔类型该写法等同于 -isbool=true)
	// -age=x		(一个 - 符号，使用等号)
	// -age x		(一个 - 符号，使用空格)
	// --age=x		(两个 - 符号，使用等号)
	// --age x		(两个 - 符号，使用空格)
	// -h/--help	可以查看帮助信息
}
