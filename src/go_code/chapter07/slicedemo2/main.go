package main
import "fmt"

//定义函数，将斐波那契数列放到切片中并返回
func fbn(n int) []uint64 {
	fbnslice := make([]uint64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i< n; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}

func main(){

	slice := fbn(5)
	fmt.Println("slice=", slice)
}