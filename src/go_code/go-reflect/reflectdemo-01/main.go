package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Score []float64 `json:"score"`
}

func main(){
	var num = 5;
	value := reflect.ValueOf(num)
	fmt.Printf("%v, %T\n", num, num)		// 5, int
	// value 为reflect.Value类型，不能与in做运算
	fmt.Printf("%v, %T\n", value, value)	// 5, reflect.Value
	fmt.Printf("%v, %T\n", value.Int(), value.Int())	// 5, int64

	if !value.CanSet(){
		fmt.Printf("v的持有值不能被修改。\n")
	}

	var stu Student = Student{"xiaoming", 20, []float64{90.2,80.5}}
	// fmt.Println(stu)
	stuType := reflect.TypeOf(stu)
	fmt.Printf("%v, %T\n", stuType, stuType)	// main.Student, *reflect.rtype
	
	// kind()返回该接口的具体分类
	fmt.Println(stuType.Kind())				// struct
	// Name返回该类型在自身包内的类型名
	fmt.Println(stuType.Name())				// Student
	// 返回struct类型的字段数
	fmt.Println(stuType.NumField())			// 3
	// 返回struct类型的第i个字段的类型,Field(i int) StructField
	// type StructField struct {
	// 	// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
	// 	// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
	// 	Name    string
	// 	PkgPath string
	// 	Type      Type      // 字段的类型
	// 	Tag       StructTag // 字段的标签
	// 	Offset    uintptr   // 字段在结构体中的字节偏移量
	// 	Index     []int     // 用于Type.FieldByIndex时的索引切片
	// 	Anonymous bool      // 是否匿名字段
	// }
	fmt.Println(stuType.Field(1))		// {Age  int json:"age" 16 [1] false}

	// 获取到reflect.Value
	rVal := reflect.ValueOf(stu)

	// 将rVal转换为interface{}类型
	iVal := rVal.Interface()

	// 将iVal转换通过断言转换为需要的类型
	va := iVal.(Student)
	fmt.Println(va)

}
