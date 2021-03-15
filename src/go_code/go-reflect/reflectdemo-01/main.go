package main

import (
	"fmt"
	"reflect"
)

// 基本数据类型
func test1(num *int){

	fmt.Printf("%v, %T\n", *num, *num)		// 5, int
	// 获取num的reflect.Value类型
	value := reflect.ValueOf(num)
	// 获取num的reflect.Type类型
	rtp := reflect.TypeOf(num)

	fmt.Printf("%v, %T\n", rtp,rtp)
	
	// value 为reflect.Value类型，不能与in做运算
	fmt.Printf("%v, %T\n", value, value)	// 5, reflect.Value

	// 得到num的接口类型
	iv := value.Interface()

	// 类型判断，转换为int
	n := iv.(*int)
	fmt.Printf("%v, %T\n", *n, *n)

	// 修改num的值
	value.Elem().SetInt(20)
}

// 结构体类型反射的基本操作
func test2(stu Student){
	// 获取到reflect.Value
	rVal := reflect.ValueOf(stu)
	// 获取到reflect.Type类型
	stuType := reflect.TypeOf(stu)

	fmt.Printf("%v, %T\n", rVal, rVal)
	fmt.Printf("%v, %T\n", stuType, stuType)	// main.Student, *reflect.rtype

	// 将rVal转换为interface{}类型
	iVal := rVal.Interface()

	// 将iVal转换通过断言转换为需要的类型
	va := iVal.(Student)
	fmt.Println(va)
	
	// kind()返回该接口的具体分类
	// fmt.Println(stuType.Kind())				// struct
	// Name返回该类型在自身包内的类型名
	// fmt.Println(stuType.Name())				// Student

	// 返回struct类型的字段数
	numField := stuType.NumField()
	// fmt.Println(numField)			// 3
	// 遍历字段
	for i:= 0; i< numField; i++ {
		fmt.Printf("第%v个字段值为%v,json串为%v\n", i+1, rVal.Field(i), stuType.Field(i).Tag.Get("json"))
	}
	// 返回struct类型的第i个字段的类型,reflect.Type类型的接口Field(i int) StructField
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

	// 返回struct的方法数,并遍历调用方法,方法必须是公有的(首字母大写),否则这里获取不到
	numMethod := rVal.NumMethod()
	fmt.Printf("方法数= %v\n", numMethod)
	for i:= 0; i< numMethod; i++{
		rVal.Method(i).Call(nil)
	}

}

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Score []float64 `json:"score"`
}

func (s Student) MyPrint(){
	fmt.Println(s.Name, s.Age, s.Score)
}

func main(){
	// 基本数据类型的反射基本操作
	var num = 5;
	test1(&num)
	fmt.Printf("修改num的值之后，num=%v\n", num)	// 20

	// 结构体类型的反射基本操作
	var stu Student = Student{"xiaoming", 20, []float64{90.2,80.5}}
	// fmt.Println(stu)
	test2(stu)
}
