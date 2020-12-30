package main
import (
	"fmt"
)

func main(){
	//数组入门

	//定义一个数组
	var hens [5]float64
	//给数组的每一个元素赋值
	hens[0] = 1.0	//数组的第一个元素
	hens[1] = 2.0	//数组的第二个元素
	hens[2] = 3.0
	hens[3] = 4.0
	hens[4] = 5.0
	//遍历数组,求出总体重
	weight := 0.0
	for i :=0; i< len(hens); i++ {
		weight += hens[i]
	}
	//求出平均体重
	aver := weight / float64(len(hens))
	fmt.Println("weight=", weight)
	fmt.Println("aver=", aver)

}