package main

import (
	"fmt"
	"reflect"
)

type Cal struct{
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string){
	res := c.Num1 - c.Num2
	fmt.Printf("%v 完成了减法运算, %v - %v = %v\n", name, c.Num1, c.Num2, res)
}

func test(i interface{}){
	rVal := reflect.ValueOf(i)
	// rTyp := reflect.TypeOf(i)

	// 遍历字段
	numFields := rVal.NumField()
	for i := 0; i < numFields; i++ {
		fmt.Printf("第%v个字段值是%v.\n", i+1, rVal.Field(i))
	}
	// 调用方法
	method := rVal.MethodByName("GetSub")
	method.Call([]reflect.Value{reflect.ValueOf("tom")})
}

func main(){
	cal := Cal{
		Num1: 20,
		Num2: 14,
	}
	test(cal)
}