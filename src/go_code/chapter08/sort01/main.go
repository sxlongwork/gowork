package main
import "fmt"

//冒泡排序
func bubbleSort(arr *[8]int) {

	fmt.Println("排序前：", *arr)

	//需要进行len(*arr)-1轮排序，每一轮比较相邻的两个数，需要比较len(*arr)-1-当前轮数
	for i := 0; i< len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			//如果前面的数比后面的数大，就交换，这是从大到小排序，
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

	fmt.Println("排序后：", *arr)
}

func main(){
	var array = [8]int{23,0,3,12,9,616,47,-8}
	bubbleSort(&array)
	fmt.Println("array=", array)
}