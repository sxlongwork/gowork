package main
import "fmt"

func test(arr *[3]int) {
	(*arr)[0] = 20
}

func main(){
	//数组是多个相同类型数据的组合，数组一旦声明/定义了，其长度是固定的，不能动态变化
	// var arr [3]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 1.2	//这里会报错，数据类型不对
	// arr[3] = 2		//这里会报错，数组越界

	var arr1 [3]float64
	var arr2 [3]string
	var arr3 [3]bool
	fmt.Printf("arr1=%v, arr2=%v, arr3=%v\n", arr1, arr2, arr3)

	//Go的数组是值类型的，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	arr := [3]int{1,2,3}
	test(&arr)
	fmt.Println("arr=",arr)
}
