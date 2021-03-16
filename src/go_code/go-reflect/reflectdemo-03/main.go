package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
}

func (s Student) MyPrint(){
	fmt.Printf("姓名：%v\t年龄：%v\t住址：%v\n", s.Name, s.Age, s.Address)
}

func (s Student) Add(n1, n2 int) {
	fmt.Printf("%v计算%v+%v? 结果=%v\n", s.Name, n1, n2, n1+n2)
}

// 结构体类型测试
func testStudent(s interface{}){

	_, ok := s.(Student)
	if !ok {
		fmt.Printf("参数类型错误, s不是结构体类型。")
		return
	}

	// 获取reflect.Value和reflect.Type
	rVal := reflect.ValueOf(s)
	rType := reflect.TypeOf(s)

	// 获取s的字段数，遍历字段并获取标签值
	numField := rVal.NumField()
	for i := 0; i< numField; i++{
		fmt.Printf("第%v个字段值为%v, 标签为%v\n", i+1, rVal.Field(i), rType.Field(i).Tag.Get("json"))
	}

	// 获取s的方法数并调用分别调用
	numMethod := rVal.NumMethod()
	fmt.Printf("该结构体共有%v个方法\n", numMethod)
	// Method方法默认按方法名排序对应的i值，所以这里调用的是MyPrint方法
	rVal.Method(1).Call(nil)

	param := make([]reflect.Value,2)
	param[0] = reflect.ValueOf(10)
	param[1] = reflect.ValueOf(20)
	rVal.Method(0).Call(param)

}

func testStudent2(s interface{}){
	rVal := reflect.ValueOf(s)
	rTyp := reflect.TypeOf(s)

	numFields := rVal.Elem().NumField()
	for i:= 0; i< numFields; i++{
		fmt.Printf("第%v个字段值为%v, 标签为%v\n", i+1, rVal.Elem().Field(i), rTyp.Elem().Field(i).Tag.Get("json"))
		fmt.Printf("修改该字段的值\n")
		filedName := rTyp.Elem().Field(i).Name
		switch filedName{
		case "Name":
			rVal.Elem().Field(i).SetString("tom")
		case "Age":
			rVal.Elem().Field(i).SetInt(3)
		case "Address":
			rVal.Elem().Field(i).SetString("beijing")
		default:
			continue
		}
	}
	// 获取s的方法数并调用分别调用
	numMethod := rVal.Elem().NumMethod()
	fmt.Printf("该结构体共有%v个方法\n", numMethod)
	// Method方法默认按方法名排序对应的i值，所以这里调用的是MyPrint方法
	rVal.Elem().Method(1).Call(nil)

	param := make([]reflect.Value,2)
	param[0] = reflect.ValueOf(100)
	param[1] = reflect.ValueOf(200)
	rVal.Elem().Method(0).Call(param)
	
}


func main(){
	var stu Student = Student{
		Name: "xiaoming",
		Age: 20,
		Address: "shanghai",
	}
	// testStudent(stu)
	testStudent2(&stu)
}