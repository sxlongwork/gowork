package main

import (
	"fmt"
	"go_code/go-struct/structdemo-11/model"
)

func main() {
	//因为student结构体首字母是小写，所以只能在model包中使用
	//要想在main包中获得struct结构体变量的信息，可以使用工厂模式
	var stu = model.GetDistance("tom", 23)
	fmt.Printf("%v的年龄是%v\n", stu.GetName(), stu.GetAge())
}
