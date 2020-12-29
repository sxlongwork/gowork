package main
import (
	"fmt"
	"math/rand"
	"time"
)

func getMax(arr *[5]int) (int, int) {
	
	index := 0
	max := (*arr)[0]
	for i := 1; i< 5; i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return index, max
}

func main(){

	var arr [26]byte
	for i := 0; i< len(arr); i++ {
		arr[i] = 'A' + byte(i)	// 'A' + 1 = 'B', 'A'参与运算会被当做byte类型，i为int类型，所以需要类型转换
	}

	// fmt.Println(arr)
	for i := 0; i< len(arr); i++ {
		fmt.Printf("%c ", arr[i])
	}
	fmt.Println()

	//求出数组的最大值
	var arr1 [5]int = [5]int{313,62,0,-4,52}
	index, max := getMax(&arr1)
	fmt.Printf("index=%d, max=%d\n", index, max)

	// 求出数组的和与平均值
	var arr2 = [...]int{11,20,30,20,10}
	sum := 0
	// for i := 0; i< 5; i++ {
	// 	sum = arr2[i] + sum
	// }
	for _, value := range arr2 {
		sum = sum + value
	}
	ave := float64(sum) / float64(len(arr2))
	fmt.Printf("sum=%v, ave=%v\n", sum, ave)

	var arr3 [5]int
	rand.Seed(time.Now().UnixNano())	//给定一个种子，保证每次执行生成的随机数不同
	for i:= 0; i< 5; i++ {
		arr3[i] = rand.Intn(100)		//生成随机数
	}
	fmt.Println("原arr3=", arr3)
	i := 0
	j := len(arr3)-1
	for ; i < j ; {
		if i < j {
			temp := arr3[i]
			arr3[i] = arr3[j]
			arr3[j] = temp
		}
		i++
		j-- 
	}
	fmt.Println("后arr3=", arr3)


}