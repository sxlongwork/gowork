package main
import "fmt"

func main(){
	// var arr [8]int = [...]int{1,2,3,4,5,6,7,8}
	// //slice是切片名称
	// slice := arr[1:4]	//注意区间是左闭右开，即包含下标为1的元素，不包含下标为4的元素
	// //for range 遍历切片
	// for index, value := range slice{
	// 	fmt.Printf("index=%v, value=%v\n", index, value)
	// }
	// fmt.Printf("slice=%v\n", slice)
	// //切片长度，切片容量
	// fmt.Printf("len(slice) = %v, cap(slice) = %v\n", len(slice),cap(slice))

	// fmt.Printf("arr[1]的地址是%p\n", &arr[1])
	// fmt.Printf("slice[0]的值是%p", &slice[0])


	var s []string = make([]string, 3)
	fmt.Printf("s = %v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))
	s[0] = "xianqian"
	s[1] = "long"
	s[2] = "love"
	fmt.Printf("s = %v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s))

	//方式3
	var slice []int = []int{34, 45, 0}
	fmt.Printf("slice=%v, len(slice)=%v, cap(slice)=%v\n", slice, len(slice), cap(slice))

	var slice2 = slice[:2]
	fmt.Printf("slice2=%v, len(slice2)=%v, cap(slice2)=%v\n", slice2, len(slice2), cap(slice2))

	var slice3 = []int{1,2,3}
	//使用append动态添加元素
	slice3 = append(slice3, 4, 5)
	fmt.Printf("slice3=%v, len(slice3)=%v, cap(slice3)=%v\n", slice3, len(slice3), cap(slice3))
	//也可以给切片追截切片
	slice3 = append(slice3, slice3...)// 也可以追加其他切片，注意...不嫩省略
	fmt.Printf("slice3=%v, len(slice3)=%v, cap(slice3)=%v\n", slice3, len(slice3), cap(slice3))

	fmt.Println("-------------------------")
	slice4 := []int{2,4,6,8}
	slice5 := make([]int, 10)
	copy(slice5, slice4)	//将slice4拷贝给slice5
	fmt.Printf("slice4=%v, len(slice4)=%v, cap(slice4)=%v\n", slice4, len(slice4), cap(slice4))
	fmt.Printf("slice5=%v, len(slice5)=%v, cap(slice5)=%v\n", slice5, len(slice5), cap(slice5))
	slice5[0] = 10	//修改slice5不会影响slice4的值，因为他们的数据空间时独立的
	fmt.Printf("slice4=%v, len(slice4)=%v, cap(slice4)=%v\n", slice4, len(slice4), cap(slice4))
	fmt.Printf("slice5=%v, len(slice5)=%v, cap(slice5)=%v\n", slice5, len(slice5), cap(slice5))

}