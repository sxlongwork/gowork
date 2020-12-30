package main

import (
	"fmt"
)


func main(){

	num1 := 100
	fmt.Printf("num1的数据类型是%T,num1=%v,num1的地址是%v\n", num1, num1, &num1)

	num2 := new(int)	//返回int类型的指针
	//num2的类型是*int
	//num2的值和地址都是由系统分配
	//要获取num2指向的值可以使用 *num2
	fmt.Printf("num2的数据类型是%T,num2=%v,num2的地址是%v，num2这个指针指向的值是%v", num2,num2,&num2,*num2)
}