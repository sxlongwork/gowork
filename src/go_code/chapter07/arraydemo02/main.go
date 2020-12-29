package main
import (
	"fmt"
)

func main(){

	// var a [3]int	//int占8个字节，默认初始值为0
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
	// fmt.Println("a = ", a)	//依次打出数组元素，若没有赋值，则打印默认初始值即3个0

	// fmt.Printf("数组a的地址是%p, a[0]的地址是%p, a[1]的地址是%p\n", &a, &a[0], &a[1])


	//数组元素的使用
	// var b [3]int
	// for i := 0; i< len(b); i++ {
	// 	fmt.Println("输入第%d个元素的值：", i+1)
	// 	fmt.Scanln(&b[i])
	// }
	// //打印数组的元素
	// for i:=0; i< len(b); i++ {
	// 	fmt.Printf("b[%d]=%d\n",i,b[i])
	// }

	//4中初始化数组的访问
	// var a1 [3]int = [3]int {1,2,3}
	// fmt.Println("a1=,", a1)
	// var a2 = [3]int{1,2,3}
	// fmt.Println("a2=,", a2)
	// var a3 = [...]int{1,2,3}
	// fmt.Println("a3=,", a3)
	// var a4 = [3]int{1:10, 0:2, 2:22}	//指定元素的下标
	// fmt.Println("a4=,", a4)

	//for-range遍历数组
	var arr [4]int = [4]int{22,12,24,45}
	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}
	//不需要index，用下划线代替
	for _, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}
}