package main

import (
	"reflect"
	"fmt"
)

type Student struct{
	Name string		`json:"name"`
	Age int			`json:"age"`
}

func test(i interface{}){
	rTyp := reflect.TypeOf(i)
	// New返回一个Value类型值，该值持有一个指向类型为typ的新申请的零值的指针
	val := reflect.New(rTyp)
	
	val.Elem().Field(0).SetString("tom")
	val.Elem().Field(1).SetInt(10)
	iv := val.Interface()
	ss := iv.(*Student)
	fmt.Println(ss.Name, ss.Age)	// tom 10

}

func main(){
	stu := Student{
		Name: "xiaoming",
		Age: 21,
	}
	test(stu)
}