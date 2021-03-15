package main

import (
	"fmt"
	"reflect"
)
func main(){
	var name = "xiaoming"
	// 获取到reflect.Value类型
	iVal := reflect.ValueOf(name)
	// 转换为interface{}
	iv := iVal.Interface()
	// 类型断言
	nv := iv.(string)
	fmt.Printf("%v, %T\n", nv, nv)	// xiaoming, string

	// 修改变量的值
	var address = "北京"
	value := reflect.ValueOf(&address)
	value.Elem().SetString("上海")
	fmt.Println(address)	// 上海

	var v float64 = 1.2
	iVal = reflect.ValueOf(v)
	fmt.Printf("%v, %v\n", iVal.Type(), iVal.Kind())
	// 将reflect.Value转换成interface{}
	iv = iVal.Interface()
	vv := iv.(float64)
	fmt.Printf("%v, %T\n", vv, vv)

	
}