package main

import (
	"fmt"
	"reflect"
)

func test1(n1, n2 int){
	fmt.Println(n1,n2)
}

func test2(n1, n2 int, str string){
	fmt.Println(n1,n2,str)
}

func test(fun interface{}, args ...interface{}){
	rVal := reflect.ValueOf(fun)
	numParam := len(args)
	param := make([]reflect.Value, numParam)
	for i:= 0; i < numParam; i++{
		param[i] = reflect.ValueOf(args[i])
	}
	rVal.Call(param)
}

func main(){
	t1 := test1
	t2 := test2
	test(t1, 2, 4)
	test(t2, 10, 20, "xiaoming")
}